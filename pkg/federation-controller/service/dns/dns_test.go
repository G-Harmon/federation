/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dns

import (
	"errors"
	"fmt"
	"sort"
	"testing"
	"time"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/federation/apis/federation/v1beta1"
	fakefedclientset "k8s.io/federation/client/clientset_generated/federation_clientset/fake"
	"k8s.io/federation/pkg/dnsprovider/providers/google/clouddns" // Only for unit testing purposes.
	"k8s.io/federation/pkg/federation-controller/service/ingress"
	. "k8s.io/federation/pkg/federation-controller/util/test"
)

const (
	cluster1Name string = "c1"
	cluster2Name string = "c2"
)

type testDeployment struct {
	name     string
	service  v1.Service
	expected sets.String
}

type NetWrapperMock struct {
	result map[string][]string
}

func (mock *NetWrapperMock) LookupHost(host string) (addrs []string, err error) {

	// If nothing to return, return empty list
	if mock.result == nil || len(mock.result) == 0 {
		return make([]string, 0), errors.New("Mock error response")
	}

	return mock.result[host], nil
}

func (mock *NetWrapperMock) AddHost(host string, addrs []string) {

	// Initialise if null
	if mock.result == nil {
		mock.result = make(map[string][]string)
	}

	mock.result[host] = addrs
}

// NewClusterWithRegionZone builds a new cluster object with given region and zone attributes.
func NewClusterWithRegionZone(name string, readyStatus v1.ConditionStatus, region, zone string) *v1beta1.Cluster {
	cluster := NewCluster(name, readyStatus)
	cluster.Status.Zones = []string{zone}
	cluster.Status.Region = region
	return cluster
}

func createTestServiceDeployments(cluster1 *v1beta1.Cluster, cluster2 *v1beta1.Cluster) []testDeployment {

	globalDNSName := "servicename.servicenamespace.myfederation.svc.federation.example.com"
	fooRegionDNSName := "servicename.servicenamespace.myfederation.svc.fooregion.federation.example.com"
	fooZoneDNSName := "servicename.servicenamespace.myfederation.svc.foozone.fooregion.federation.example.com"
	barRegionDNSName := "servicename.servicenamespace.myfederation.svc.barregion.federation.example.com"
	barZoneDNSName := "servicename.servicenamespace.myfederation.svc.barzone.barregion.federation.example.com"

	tests := []testDeployment{
		{
			name: "ServiceWithSingleLBIngress",
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{
					ingress.FederatedServiceIngressAnnotation: ingress.NewFederatedServiceIngress().
						AddEndpoints(cluster1Name, []string{"198.51.100.1"}).
						AddEndpoints(cluster2Name, []string{}).
						String()},
				},
			},
			expected: sets.NewString(
				"example.com:"+globalDNSName+":A:180:[198.51.100.1]",
				"example.com:"+fooRegionDNSName+":A:180:[198.51.100.1]",
				"example.com:"+fooZoneDNSName+":A:180:[198.51.100.1]",
				"example.com:"+barRegionDNSName+":CNAME:180:["+globalDNSName+"]",
				"example.com:"+barZoneDNSName+":CNAME:180:["+barRegionDNSName+"]",
			),
		},
		{
			name: "ServiceWithNoLBIngress",
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{
					ingress.FederatedServiceIngressAnnotation: ingress.NewFederatedServiceIngress().
						AddEndpoints(cluster1Name, []string{}).
						AddEndpoints(cluster2Name, []string{}).
						String()},
				},
			},
			expected: sets.NewString(
				"example.com:"+fooRegionDNSName+":CNAME:180:["+globalDNSName+"]",
				"example.com:"+fooZoneDNSName+":CNAME:180:["+fooRegionDNSName+"]",
				"example.com:"+barRegionDNSName+":CNAME:180:["+globalDNSName+"]",
				"example.com:"+barZoneDNSName+":CNAME:180:["+barRegionDNSName+"]",
			),
		},
		{
			name: "ServiceWithMultipleLBIngress",
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{
					ingress.FederatedServiceIngressAnnotation: ingress.NewFederatedServiceIngress().
						AddEndpoints(cluster1Name, []string{"198.51.100.1"}).
						AddEndpoints(cluster2Name, []string{"198.51.200.1"}).
						String()},
				},
			},
			expected: sets.NewString(
				"example.com:"+globalDNSName+":A:180:[198.51.100.1 198.51.200.1]",
				"example.com:"+fooRegionDNSName+":A:180:[198.51.100.1]",
				"example.com:"+fooZoneDNSName+":A:180:[198.51.100.1]",
				"example.com:"+barRegionDNSName+":A:180:[198.51.200.1]",
				"example.com:"+barZoneDNSName+":A:180:[198.51.200.1]",
			),
		},
		{
			// Tests getResolvedEndpoints DNS lookup
			name: "ServiceWithAWSAndGCPMultipleLBIngress",
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{
					ingress.FederatedServiceIngressAnnotation: ingress.NewFederatedServiceIngress().
						AddHostnameEndpoints(cluster1Name, []string{"a9.us-west-2.elb.amazonaws.com"}).
						AddEndpoints(cluster2Name, []string{"198.51.200.1"}).
						String()},
				},
			},
			expected: sets.NewString(
				"example.com:"+globalDNSName+":A:180:[198.51.100.1 198.51.200.1]",
				"example.com:"+fooRegionDNSName+":A:180:[198.51.100.1]",
				"example.com:"+fooZoneDNSName+":A:180:[198.51.100.1]",
				"example.com:"+barRegionDNSName+":A:180:[198.51.200.1]",
				"example.com:"+barZoneDNSName+":A:180:[198.51.200.1]",
			),
		},
		{
			name: "ServiceWithLBIngressAndServiceDeleted",
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{
					ingress.FederatedServiceIngressAnnotation: ingress.NewFederatedServiceIngress().
						AddEndpoints(cluster1Name, []string{"198.51.100.1"}).
						AddEndpoints(cluster2Name, []string{"198.51.200.1"}).
						String()},
					DeletionTimestamp: &metav1.Time{Time: time.Now()},
				},
			},
			expected: sets.NewString(
				// TODO: Ideally we should expect that there are no DNS records
				// when federated service is deleted. Need to remove these leaks in future
				"example.com:"+fooRegionDNSName+":CNAME:180:["+globalDNSName+"]",
				"example.com:"+fooZoneDNSName+":CNAME:180:["+fooRegionDNSName+"]",
				"example.com:"+barRegionDNSName+":CNAME:180:["+globalDNSName+"]",
				"example.com:"+barZoneDNSName+":CNAME:180:["+barRegionDNSName+"]",
			),
		},
		{
			name: "ServiceWithMultipleLBIngressAndOneLBIngressGettingRemoved",
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{
					ingress.FederatedServiceIngressAnnotation: ingress.NewFederatedServiceIngress().
						AddEndpoints(cluster1Name, []string{"198.51.100.1"}).
						AddEndpoints(cluster2Name, []string{"198.51.200.1"}).
						RemoveEndpoint(cluster2Name, "198.51.200.1").
						String()},
				},
			},
			expected: sets.NewString(
				"example.com:"+globalDNSName+":A:180:[198.51.100.1]",
				"example.com:"+fooRegionDNSName+":A:180:[198.51.100.1]",
				"example.com:"+fooZoneDNSName+":A:180:[198.51.100.1]",
				"example.com:"+barRegionDNSName+":CNAME:180:["+globalDNSName+"]",
				"example.com:"+barZoneDNSName+":CNAME:180:["+barRegionDNSName+"]",
			),
		},
		{
			name: "ServiceWithMultipleLBIngressAndAllLBIngressGettingRemoved",
			service: v1.Service{
				ObjectMeta: metav1.ObjectMeta{Annotations: map[string]string{
					ingress.FederatedServiceIngressAnnotation: ingress.NewFederatedServiceIngress().
						AddEndpoints(cluster1Name, []string{"198.51.100.1"}).
						AddEndpoints(cluster2Name, []string{"198.51.200.1"}).
						RemoveEndpoint(cluster1Name, "198.51.100.1").
						RemoveEndpoint(cluster2Name, "198.51.200.1").
						String()},
				},
			},
			expected: sets.NewString(
				"example.com:"+fooRegionDNSName+":CNAME:180:["+globalDNSName+"]",
				"example.com:"+fooZoneDNSName+":CNAME:180:["+fooRegionDNSName+"]",
				"example.com:"+barRegionDNSName+":CNAME:180:["+globalDNSName+"]",
				"example.com:"+barZoneDNSName+":CNAME:180:["+barRegionDNSName+"]",
			),
		},
	}
	return tests
}

func TestServiceController_ensureDnsRecords(t *testing.T) {

	cluster1 := NewClusterWithRegionZone(cluster1Name, v1.ConditionTrue, "fooregion", "foozone")
	cluster2 := NewClusterWithRegionZone(cluster2Name, v1.ConditionTrue, "barregion", "barzone")

	tests := createTestServiceDeployments(cluster1, cluster2)

	for _, test := range tests {

		fmt.Println("Running test: ", test.name)

		fakedns, _ := clouddns.NewFakeInterface()
		fakednsZones, ok := fakedns.Zones()
		if !ok {
			t.Error("Unable to fetch zones")
		}

		fakeClient := &fakefedclientset.Clientset{}
		RegisterFakeClusterGet(&fakeClient.Fake, &v1beta1.ClusterList{Items: []v1beta1.Cluster{*cluster1, *cluster2}})

		// Fake out the internet
		netmock := &NetWrapperMock{}
		netmock.AddHost("a9.us-west-2.elb.amazonaws.com", []string{"198.51.100.1"})

		d := ServiceDNSController{
			federationClient: fakeClient,
			dns:              fakedns,
			dnsZones:         fakednsZones,
			serviceDNSSuffix: "federation.example.com",
			zoneName:         "example.com",
			federationName:   "myfederation",
			netWrapper:       netmock,
		}

		dnsZones, err := getDNSZones(d.zoneName, d.zoneID, d.dnsZones)
		if err != nil {
			t.Errorf("Test failed for %s, Get DNS Zones failed: %v", test.name, err)
		}
		d.dnsZone = dnsZones[0]
		test.service.Name = "servicename"
		test.service.Namespace = "servicenamespace"

		ingress, err := ingress.ParseFederatedServiceIngress(&test.service)
		if err != nil {
			t.Errorf("Error in parsing lb ingress for service %s/%s: %v", test.service.Namespace, test.service.Name, err)
			return
		}
		for _, clusterIngress := range ingress.Items {
			// **** Main function being tested ****
			d.ensureDNSRecords(clusterIngress.Cluster, &test.service)
		}

		zones, err := fakednsZones.List()
		if err != nil {
			t.Errorf("error querying zones: %v", err)
		}

		// Dump every record to a testable-by-string-comparison form
		records := sets.NewString()
		for _, z := range zones {
			zoneName := z.Name()

			rrs, ok := z.ResourceRecordSets()
			if !ok {
				t.Errorf("cannot get rrs for zone %q", zoneName)
			}

			rrList, err := rrs.List()
			if err != nil {
				t.Errorf("error querying rr for zone %q: %v", zoneName, err)
			}
			for _, rr := range rrList {
				rrdatas := rr.Rrdatas()

				// Put in consistent (testable-by-string-comparison) order
				sort.Strings(rrdatas)
				records.Insert(fmt.Sprintf("%s:%s:%s:%d:%s", zoneName, rr.Name(), rr.Type(), rr.Ttl(), rrdatas))
			}
		}

		if !records.Equal(test.expected) {
			t.Errorf("Test %q failed.  Actual=%v, Expected=%v", test.name, records, test.expected)
		}
	}
}
