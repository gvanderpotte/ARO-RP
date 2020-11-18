# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft and contributors.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#
# See the License for the specific language governing permissions and
# limitations under the License.
#
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.
# --------------------------------------------------------------------------

from msrest.serialization import Model
from msrest.exceptions import HttpOperationError


class APIServerProfile(Model):
    """APIServerProfile represents an API server profile.

    :param visibility: API server visibility (immutable). Possible values
     include: 'Private', 'Public'
    :type visibility: str or
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.enum
    :param url: The URL to access the cluster API server (immutable).
    :type url: str
    :param ip: The IP of the cluster API server (immutable).
    :type ip: str
    """

    _attribute_map = {
        'visibility': {'key': 'visibility', 'type': 'str'},
        'url': {'key': 'url', 'type': 'str'},
        'ip': {'key': 'ip', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(APIServerProfile, self).__init__(**kwargs)
        self.visibility = kwargs.get('visibility', None)
        self.url = kwargs.get('url', None)
        self.ip = kwargs.get('ip', None)


class Resource(Model):
    """Resource.

    Variables are only populated by the server, and will be ignored when
    sending a request.

    :ivar id: Fully qualified resource Id for the resource. Ex -
     /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
    :vartype id: str
    :ivar name: The name of the resource
    :vartype name: str
    :ivar type: The type of the resource. Ex-
     Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    :vartype type: str
    """

    _validation = {
        'id': {'readonly': True},
        'name': {'readonly': True},
        'type': {'readonly': True},
    }

    _attribute_map = {
        'id': {'key': 'id', 'type': 'str'},
        'name': {'key': 'name', 'type': 'str'},
        'type': {'key': 'type', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(Resource, self).__init__(**kwargs)
        self.id = None
        self.name = None
        self.type = None


class AzureEntityResource(Resource):
    """The resource model definition for a Azure Resource Manager resource with an
    etag.

    Variables are only populated by the server, and will be ignored when
    sending a request.

    :ivar id: Fully qualified resource Id for the resource. Ex -
     /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
    :vartype id: str
    :ivar name: The name of the resource
    :vartype name: str
    :ivar type: The type of the resource. Ex-
     Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    :vartype type: str
    :ivar etag: Resource Etag.
    :vartype etag: str
    """

    _validation = {
        'id': {'readonly': True},
        'name': {'readonly': True},
        'type': {'readonly': True},
        'etag': {'readonly': True},
    }

    _attribute_map = {
        'id': {'key': 'id', 'type': 'str'},
        'name': {'key': 'name', 'type': 'str'},
        'type': {'key': 'type', 'type': 'str'},
        'etag': {'key': 'etag', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(AzureEntityResource, self).__init__(**kwargs)
        self.etag = None


class CloudError(Model):
    """CloudError represents a cloud error.

    :param error: An error response from the service.
    :type error: ~azure.mgmt.redhatopenshift.v2020_04_30.models.CloudErrorBody
    """

    _attribute_map = {
        'error': {'key': 'error', 'type': 'CloudErrorBody'},
    }

    def __init__(self, **kwargs):
        super(CloudError, self).__init__(**kwargs)
        self.error = kwargs.get('error', None)


class CloudErrorException(HttpOperationError):
    """Server responsed with exception of type: 'CloudError'.

    :param deserialize: A deserializer
    :param response: Server response to be deserialized.
    """

    def __init__(self, deserialize, response, *args):

        super(CloudErrorException, self).__init__(deserialize, response, 'CloudError', *args)


class CloudErrorBody(Model):
    """CloudErrorBody represents the body of a cloud error.

    :param code: An identifier for the error. Codes are invariant and are
     intended to be consumed programmatically.
    :type code: str
    :param message: A message describing the error, intended to be suitable
     for display in a user interface.
    :type message: str
    :param target: The target of the particular error. For example, the name
     of the property in error.
    :type target: str
    :param details: A list of additional details about the error.
    :type details:
     list[~azure.mgmt.redhatopenshift.v2020_04_30.models.CloudErrorBody]
    """

    _attribute_map = {
        'code': {'key': 'code', 'type': 'str'},
        'message': {'key': 'message', 'type': 'str'},
        'target': {'key': 'target', 'type': 'str'},
        'details': {'key': 'details', 'type': '[CloudErrorBody]'},
    }

    def __init__(self, **kwargs):
        super(CloudErrorBody, self).__init__(**kwargs)
        self.code = kwargs.get('code', None)
        self.message = kwargs.get('message', None)
        self.target = kwargs.get('target', None)
        self.details = kwargs.get('details', None)


class ClusterProfile(Model):
    """ClusterProfile represents a cluster profile.

    :param pull_secret: The pull secret for the cluster (immutable).
    :type pull_secret: str
    :param domain: The domain for the cluster (immutable).
    :type domain: str
    :param version: The version of the cluster (immutable).
    :type version: str
    :param resource_group_id: The ID of the cluster resource group
     (immutable).
    :type resource_group_id: str
    """

    _attribute_map = {
        'pull_secret': {'key': 'pullSecret', 'type': 'str'},
        'domain': {'key': 'domain', 'type': 'str'},
        'version': {'key': 'version', 'type': 'str'},
        'resource_group_id': {'key': 'resourceGroupId', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(ClusterProfile, self).__init__(**kwargs)
        self.pull_secret = kwargs.get('pull_secret', None)
        self.domain = kwargs.get('domain', None)
        self.version = kwargs.get('version', None)
        self.resource_group_id = kwargs.get('resource_group_id', None)


class ConsoleProfile(Model):
    """ConsoleProfile represents a console profile.

    :param url: The URL to access the cluster console (immutable).
    :type url: str
    """

    _attribute_map = {
        'url': {'key': 'url', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(ConsoleProfile, self).__init__(**kwargs)
        self.url = kwargs.get('url', None)


class Display(Model):
    """Display represents the display details of an operation.

    :param provider: Friendly name of the resource provider.
    :type provider: str
    :param resource: Resource type on which the operation is performed.
    :type resource: str
    :param operation: Operation type: read, write, delete, listKeys/action,
     etc.
    :type operation: str
    :param description: Friendly name of the operation.
    :type description: str
    """

    _attribute_map = {
        'provider': {'key': 'provider', 'type': 'str'},
        'resource': {'key': 'resource', 'type': 'str'},
        'operation': {'key': 'operation', 'type': 'str'},
        'description': {'key': 'description', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(Display, self).__init__(**kwargs)
        self.provider = kwargs.get('provider', None)
        self.resource = kwargs.get('resource', None)
        self.operation = kwargs.get('operation', None)
        self.description = kwargs.get('description', None)


class IngressProfile(Model):
    """IngressProfile represents an ingress profile.

    :param name: The ingress profile name.  Must be "default" (immutable).
    :type name: str
    :param visibility: Ingress visibility (immutable). Possible values
     include: 'Private', 'Public'
    :type visibility: str or
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.enum
    :param ip: The IP of the ingress (immutable).
    :type ip: str
    """

    _attribute_map = {
        'name': {'key': 'name', 'type': 'str'},
        'visibility': {'key': 'visibility', 'type': 'str'},
        'ip': {'key': 'ip', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(IngressProfile, self).__init__(**kwargs)
        self.name = kwargs.get('name', None)
        self.visibility = kwargs.get('visibility', None)
        self.ip = kwargs.get('ip', None)


class MasterProfile(Model):
    """MasterProfile represents a master profile.

    :param vm_size: The size of the master VMs (immutable). Possible values
     include: 'Standard_D16as_v4', 'Standard_D16s_v3', 'Standard_D2s_v3',
     'Standard_D32as_v4', 'Standard_D32s_v3', 'Standard_D4as_v4',
     'Standard_D4s_v3', 'Standard_D8as_v4', 'Standard_D8s_v3',
     'Standard_E16s_v3', 'Standard_E32s_v3', 'Standard_E4s_v3',
     'Standard_E8s_v3', 'Standard_F16s_v2', 'Standard_F32s_v2',
     'Standard_F4s_v2', 'Standard_F8s_v2'
    :type vm_size: str or ~azure.mgmt.redhatopenshift.v2020_04_30.models.enum
    :param subnet_id: The Azure resource ID of the master subnet (immutable).
    :type subnet_id: str
    """

    _attribute_map = {
        'vm_size': {'key': 'vmSize', 'type': 'str'},
        'subnet_id': {'key': 'subnetId', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(MasterProfile, self).__init__(**kwargs)
        self.vm_size = kwargs.get('vm_size', None)
        self.subnet_id = kwargs.get('subnet_id', None)


class NetworkProfile(Model):
    """NetworkProfile represents a network profile.

    :param pod_cidr: The CIDR used for OpenShift/Kubernetes Pods (immutable).
    :type pod_cidr: str
    :param service_cidr: The CIDR used for OpenShift/Kubernetes Services
     (immutable).
    :type service_cidr: str
    """

    _attribute_map = {
        'pod_cidr': {'key': 'podCidr', 'type': 'str'},
        'service_cidr': {'key': 'serviceCidr', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(NetworkProfile, self).__init__(**kwargs)
        self.pod_cidr = kwargs.get('pod_cidr', None)
        self.service_cidr = kwargs.get('service_cidr', None)


class TrackedResource(Resource):
    """The resource model definition for a ARM tracked top level resource.

    Variables are only populated by the server, and will be ignored when
    sending a request.

    All required parameters must be populated in order to send to Azure.

    :ivar id: Fully qualified resource Id for the resource. Ex -
     /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
    :vartype id: str
    :ivar name: The name of the resource
    :vartype name: str
    :ivar type: The type of the resource. Ex-
     Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    :vartype type: str
    :param tags: Resource tags.
    :type tags: dict[str, str]
    :param location: Required. The geo-location where the resource lives
    :type location: str
    """

    _validation = {
        'id': {'readonly': True},
        'name': {'readonly': True},
        'type': {'readonly': True},
        'location': {'required': True},
    }

    _attribute_map = {
        'id': {'key': 'id', 'type': 'str'},
        'name': {'key': 'name', 'type': 'str'},
        'type': {'key': 'type', 'type': 'str'},
        'tags': {'key': 'tags', 'type': '{str}'},
        'location': {'key': 'location', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(TrackedResource, self).__init__(**kwargs)
        self.tags = kwargs.get('tags', None)
        self.location = kwargs.get('location', None)


class OpenShiftCluster(TrackedResource):
    """OpenShiftCluster represents an Azure Red Hat OpenShift cluster.

    Variables are only populated by the server, and will be ignored when
    sending a request.

    All required parameters must be populated in order to send to Azure.

    :ivar id: Fully qualified resource Id for the resource. Ex -
     /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
    :vartype id: str
    :ivar name: The name of the resource
    :vartype name: str
    :ivar type: The type of the resource. Ex-
     Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    :vartype type: str
    :param tags: Resource tags.
    :type tags: dict[str, str]
    :param location: Required. The geo-location where the resource lives
    :type location: str
    :param provisioning_state: The cluster provisioning state (immutable).
     Possible values include: 'AdminUpdating', 'Creating', 'Deleting',
     'Failed', 'Succeeded', 'Updating'
    :type provisioning_state: str or
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.enum
    :param cluster_profile: The cluster profile.
    :type cluster_profile:
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.ClusterProfile
    :param console_profile: The console profile.
    :type console_profile:
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.ConsoleProfile
    :param service_principal_profile: The cluster service principal profile.
    :type service_principal_profile:
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.ServicePrincipalProfile
    :param network_profile: The cluster network profile.
    :type network_profile:
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.NetworkProfile
    :param master_profile: The cluster master profile.
    :type master_profile:
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.MasterProfile
    :param worker_profiles: The cluster worker profiles.
    :type worker_profiles:
     list[~azure.mgmt.redhatopenshift.v2020_04_30.models.WorkerProfile]
    :param apiserver_profile: The cluster API server profile.
    :type apiserver_profile:
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.APIServerProfile
    :param ingress_profiles: The cluster ingress profiles.
    :type ingress_profiles:
     list[~azure.mgmt.redhatopenshift.v2020_04_30.models.IngressProfile]
    """

    _validation = {
        'id': {'readonly': True},
        'name': {'readonly': True},
        'type': {'readonly': True},
        'location': {'required': True},
    }

    _attribute_map = {
        'id': {'key': 'id', 'type': 'str'},
        'name': {'key': 'name', 'type': 'str'},
        'type': {'key': 'type', 'type': 'str'},
        'tags': {'key': 'tags', 'type': '{str}'},
        'location': {'key': 'location', 'type': 'str'},
        'provisioning_state': {'key': 'properties.provisioningState', 'type': 'str'},
        'cluster_profile': {'key': 'properties.clusterProfile', 'type': 'ClusterProfile'},
        'console_profile': {'key': 'properties.consoleProfile', 'type': 'ConsoleProfile'},
        'service_principal_profile': {'key': 'properties.servicePrincipalProfile', 'type': 'ServicePrincipalProfile'},
        'network_profile': {'key': 'properties.networkProfile', 'type': 'NetworkProfile'},
        'master_profile': {'key': 'properties.masterProfile', 'type': 'MasterProfile'},
        'worker_profiles': {'key': 'properties.workerProfiles', 'type': '[WorkerProfile]'},
        'apiserver_profile': {'key': 'properties.apiserverProfile', 'type': 'APIServerProfile'},
        'ingress_profiles': {'key': 'properties.ingressProfiles', 'type': '[IngressProfile]'},
    }

    def __init__(self, **kwargs):
        super(OpenShiftCluster, self).__init__(**kwargs)
        self.provisioning_state = kwargs.get('provisioning_state', None)
        self.cluster_profile = kwargs.get('cluster_profile', None)
        self.console_profile = kwargs.get('console_profile', None)
        self.service_principal_profile = kwargs.get('service_principal_profile', None)
        self.network_profile = kwargs.get('network_profile', None)
        self.master_profile = kwargs.get('master_profile', None)
        self.worker_profiles = kwargs.get('worker_profiles', None)
        self.apiserver_profile = kwargs.get('apiserver_profile', None)
        self.ingress_profiles = kwargs.get('ingress_profiles', None)


class OpenShiftClusterCredentials(Model):
    """OpenShiftClusterCredentials represents an OpenShift cluster's credentials.

    :param kubeadmin_username: The username for the kubeadmin user
    :type kubeadmin_username: str
    :param kubeadmin_password: The password for the kubeadmin user
    :type kubeadmin_password: str
    """

    _attribute_map = {
        'kubeadmin_username': {'key': 'kubeadminUsername', 'type': 'str'},
        'kubeadmin_password': {'key': 'kubeadminPassword', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(OpenShiftClusterCredentials, self).__init__(**kwargs)
        self.kubeadmin_username = kwargs.get('kubeadmin_username', None)
        self.kubeadmin_password = kwargs.get('kubeadmin_password', None)


class OpenShiftClusterUpdate(Model):
    """OpenShiftCluster represents an Azure Red Hat OpenShift cluster.

    :param tags: The resource tags.
    :type tags: dict[str, str]
    :param provisioning_state: The cluster provisioning state (immutable).
     Possible values include: 'AdminUpdating', 'Creating', 'Deleting',
     'Failed', 'Succeeded', 'Updating'
    :type provisioning_state: str or
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.enum
    :param cluster_profile: The cluster profile.
    :type cluster_profile:
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.ClusterProfile
    :param console_profile: The console profile.
    :type console_profile:
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.ConsoleProfile
    :param service_principal_profile: The cluster service principal profile.
    :type service_principal_profile:
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.ServicePrincipalProfile
    :param network_profile: The cluster network profile.
    :type network_profile:
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.NetworkProfile
    :param master_profile: The cluster master profile.
    :type master_profile:
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.MasterProfile
    :param worker_profiles: The cluster worker profiles.
    :type worker_profiles:
     list[~azure.mgmt.redhatopenshift.v2020_04_30.models.WorkerProfile]
    :param apiserver_profile: The cluster API server profile.
    :type apiserver_profile:
     ~azure.mgmt.redhatopenshift.v2020_04_30.models.APIServerProfile
    :param ingress_profiles: The cluster ingress profiles.
    :type ingress_profiles:
     list[~azure.mgmt.redhatopenshift.v2020_04_30.models.IngressProfile]
    """

    _attribute_map = {
        'tags': {'key': 'tags', 'type': '{str}'},
        'provisioning_state': {'key': 'properties.provisioningState', 'type': 'str'},
        'cluster_profile': {'key': 'properties.clusterProfile', 'type': 'ClusterProfile'},
        'console_profile': {'key': 'properties.consoleProfile', 'type': 'ConsoleProfile'},
        'service_principal_profile': {'key': 'properties.servicePrincipalProfile', 'type': 'ServicePrincipalProfile'},
        'network_profile': {'key': 'properties.networkProfile', 'type': 'NetworkProfile'},
        'master_profile': {'key': 'properties.masterProfile', 'type': 'MasterProfile'},
        'worker_profiles': {'key': 'properties.workerProfiles', 'type': '[WorkerProfile]'},
        'apiserver_profile': {'key': 'properties.apiserverProfile', 'type': 'APIServerProfile'},
        'ingress_profiles': {'key': 'properties.ingressProfiles', 'type': '[IngressProfile]'},
    }

    def __init__(self, **kwargs):
        super(OpenShiftClusterUpdate, self).__init__(**kwargs)
        self.tags = kwargs.get('tags', None)
        self.provisioning_state = kwargs.get('provisioning_state', None)
        self.cluster_profile = kwargs.get('cluster_profile', None)
        self.console_profile = kwargs.get('console_profile', None)
        self.service_principal_profile = kwargs.get('service_principal_profile', None)
        self.network_profile = kwargs.get('network_profile', None)
        self.master_profile = kwargs.get('master_profile', None)
        self.worker_profiles = kwargs.get('worker_profiles', None)
        self.apiserver_profile = kwargs.get('apiserver_profile', None)
        self.ingress_profiles = kwargs.get('ingress_profiles', None)


class Operation(Model):
    """Operation represents an RP operation.

    :param name: Operation name: {provider}/{resource}/{operation}.
    :type name: str
    :param display: The object that describes the operation.
    :type display: ~azure.mgmt.redhatopenshift.v2020_04_30.models.Display
    :param origin: Sources of requests to this operation.  Comma separated
     list with valid values user or system, e.g. "user,system".
    :type origin: str
    """

    _attribute_map = {
        'name': {'key': 'name', 'type': 'str'},
        'display': {'key': 'display', 'type': 'Display'},
        'origin': {'key': 'origin', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(Operation, self).__init__(**kwargs)
        self.name = kwargs.get('name', None)
        self.display = kwargs.get('display', None)
        self.origin = kwargs.get('origin', None)


class ProxyResource(Resource):
    """The resource model definition for a ARM proxy resource. It will have
    everything other than required location and tags.

    Variables are only populated by the server, and will be ignored when
    sending a request.

    :ivar id: Fully qualified resource Id for the resource. Ex -
     /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
    :vartype id: str
    :ivar name: The name of the resource
    :vartype name: str
    :ivar type: The type of the resource. Ex-
     Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    :vartype type: str
    """

    _validation = {
        'id': {'readonly': True},
        'name': {'readonly': True},
        'type': {'readonly': True},
    }

    _attribute_map = {
        'id': {'key': 'id', 'type': 'str'},
        'name': {'key': 'name', 'type': 'str'},
        'type': {'key': 'type', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(ProxyResource, self).__init__(**kwargs)


class ServicePrincipalProfile(Model):
    """ServicePrincipalProfile represents a service principal profile.

    :param client_id: The client ID used for the cluster (immutable).
    :type client_id: str
    :param client_secret: The client secret used for the cluster (immutable).
    :type client_secret: str
    """

    _attribute_map = {
        'client_id': {'key': 'clientId', 'type': 'str'},
        'client_secret': {'key': 'clientSecret', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(ServicePrincipalProfile, self).__init__(**kwargs)
        self.client_id = kwargs.get('client_id', None)
        self.client_secret = kwargs.get('client_secret', None)


class WorkerProfile(Model):
    """WorkerProfile represents a worker profile.

    :param name: The worker profile name.  Must be "worker" (immutable).
    :type name: str
    :param vm_size: The size of the worker VMs (immutable). Possible values
     include: 'Standard_D16as_v4', 'Standard_D16s_v3', 'Standard_D2s_v3',
     'Standard_D32as_v4', 'Standard_D32s_v3', 'Standard_D4as_v4',
     'Standard_D4s_v3', 'Standard_D8as_v4', 'Standard_D8s_v3',
     'Standard_E16s_v3', 'Standard_E32s_v3', 'Standard_E4s_v3',
     'Standard_E8s_v3', 'Standard_F16s_v2', 'Standard_F32s_v2',
     'Standard_F4s_v2', 'Standard_F8s_v2'
    :type vm_size: str or ~azure.mgmt.redhatopenshift.v2020_04_30.models.enum
    :param disk_size_gb: The disk size of the worker VMs.  Must be 128 or
     greater (immutable).
    :type disk_size_gb: int
    :param subnet_id: The Azure resource ID of the worker subnet (immutable).
    :type subnet_id: str
    :param count: The number of worker VMs.
    :type count: int
    """

    _attribute_map = {
        'name': {'key': 'name', 'type': 'str'},
        'vm_size': {'key': 'vmSize', 'type': 'str'},
        'disk_size_gb': {'key': 'diskSizeGB', 'type': 'int'},
        'subnet_id': {'key': 'subnetId', 'type': 'str'},
        'count': {'key': 'count', 'type': 'int'},
    }

    def __init__(self, **kwargs):
        super(WorkerProfile, self).__init__(**kwargs)
        self.name = kwargs.get('name', None)
        self.vm_size = kwargs.get('vm_size', None)
        self.disk_size_gb = kwargs.get('disk_size_gb', None)
        self.subnet_id = kwargs.get('subnet_id', None)
        self.count = kwargs.get('count', None)
