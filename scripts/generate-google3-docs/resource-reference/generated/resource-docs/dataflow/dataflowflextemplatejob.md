{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}


{% block page_title %}DataflowFlexTemplateJob{% endblock %}
{% block body %}



Note: DataflowFlexTemplateJob resources  cannot be updated.

<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>Cloud Dataflow</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/dataflow/docs/">/dataflow/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>v1b3.projects.jobs</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/dataflow/docs/reference/rest/v1b3/projects.jobs">/dataflow/docs/reference/rest/v1b3/projects.jobs</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpdataflowflextemplatejob<br>gcpdataflowflextemplatejobs<br>dataflowflextemplatejob</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>dataflow.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>dataflowflextemplatejobs.dataflow.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>No</td>
</tr>


</tbody>
</table>

## Custom Resource Definition Properties


### Annotations
<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>cnrm.cloud.google.com/on-delete</code></td>
    </tr>
    <tr>
        <td><code>cnrm.cloud.google.com/project-id</code></td>
    </tr>
    <tr>
        <td><code>cnrm.cloud.google.com/skip-wait-on-job-termination</code></td>
    </tr>
</tbody>
</table>


### Spec
#### Schema
  ```yaml
  containerSpecGcsPath: string
  parameters: {}
  region: string
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>containerSpecGcsPath</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>parameters</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>region</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The region in which the created job should run.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>



### Status
#### Schema
  ```yaml
  conditions:
  - lastTransitionTime: string
    message: string
    reason: string
    status: string
    type: string
  jobId: string
  observedGeneration: integer
  state: string
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observation of the resource's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>jobId</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>state</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Batch Dataflow Flex Template Job
  ```yaml
  # Copyright 2020 Google LLC
  #
  # Licensed under the Apache License, Version 2.0 (the "License");
  # you may not use this file except in compliance with the License.
  # You may obtain a copy of the License at
  #
  #     http://www.apache.org/licenses/LICENSE-2.0
  #
  # Unless required by applicable law or agreed to in writing, software
  # distributed under the License is distributed on an "AS IS" BASIS,
  # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  # See the License for the specific language governing permissions and
  # limitations under the License.
  
  apiVersion: dataflow.cnrm.cloud.google.com/v1beta1
  kind: DataflowFlexTemplateJob
  metadata:
    annotations:
      cnrm.cloud.google.com/on-delete: "cancel"
    name: dataflowflextemplatejob-sample-batch
  spec:
    region: us-central1
    # This is a public, Google-maintained Dataflow Job flex template of a batch job
    containerSpecGcsPath: gs://dataflow-templates/2020-08-31-00_RC00/flex/File_Format_Conversion
    parameters:
      inputFileFormat: csv
      outputFileFormat: avro
      # This is a public, Google-maintained csv file expressly for this sample.
      inputFileSpec: gs://config-connector-samples/dataflowflextemplate/numbertest.csv
      # Replace ${PROJECT_ID?} with your project ID.
      outputBucket: gs://${PROJECT_ID?}-dataflowflextemplatejob-dep-batch
      # This is a public, Google-maintained Avro schema file expressly for this sample.
      schema: gs://config-connector-samples/dataflowflextemplate/numbers.avsc
  ---
  apiVersion: storage.cnrm.cloud.google.com/v1beta1
  kind: StorageBucket
  metadata:
    # StorageBucket names must be globally unique. Replace ${PROJECT_ID?} with your project ID.
    name: ${PROJECT_ID?}-dataflowflextemplatejob-dep-batch
  ```

### Streaming Dataflow Flex Template Job
  ```yaml
  # Copyright 2020 Google LLC
  #
  # Licensed under the Apache License, Version 2.0 (the "License");
  # you may not use this file except in compliance with the License.
  # You may obtain a copy of the License at
  #
  #     http://www.apache.org/licenses/LICENSE-2.0
  #
  # Unless required by applicable law or agreed to in writing, software
  # distributed under the License is distributed on an "AS IS" BASIS,
  # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  # See the License for the specific language governing permissions and
  # limitations under the License.
  
  apiVersion: dataflow.cnrm.cloud.google.com/v1beta1
  kind: DataflowFlexTemplateJob
  metadata:
    annotations:
      cnrm.cloud.google.com/on-delete: "drain"
    name: dataflowflextemplatejob-sample-streaming
  spec:
    region: us-central1
    # This is a public, Google-maintained Dataflow Job flex template of a streaming job
    containerSpecGcsPath: gs://dataflow-templates/2020-08-31-00_RC00/flex/PubSub_Avro_to_BigQuery
    parameters:
      # This is a public, Google-maintained Avro schema file expressly for this sample.
      schemaPath: gs://config-connector-samples/dataflowflextemplate/numbers.avsc
      # Replace ${PROJECT_ID?} with your project ID.
      inputSubscription: projects/${PROJECT_ID?}/subscriptions/dataflowflextemplatejob-dep-streaming
      outputTopic: projects/${PROJECT_ID?}/topics/dataflowflextemplatejob-dep1-streaming
      outputTableSpec: ${PROJECT_ID?}:dataflowflextemplatejobdepstreaming.dataflowflextemplatejobdepstreaming
      createDisposition: CREATE_NEVER
  ---
  apiVersion: bigquery.cnrm.cloud.google.com/v1beta1
  kind: BigQueryDataset
  metadata:
    name: dataflowflextemplatejobdepstreaming
  ---
  apiVersion: bigquery.cnrm.cloud.google.com/v1beta1
  kind: BigQueryTable
  metadata:
    name: dataflowflextemplatejobdepstreaming
  spec:
    datasetRef:
      name: dataflowflextemplatejobdepstreaming
  ---
  apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
  kind: PubSubSubscription
  metadata:
    name: dataflowflextemplatejob-dep-streaming
  spec:
    topicRef:
      name: dataflowflextemplatejob-dep0-streaming
  ---
  apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
  kind: PubSubTopic
  metadata:
    name: dataflowflextemplatejob-dep0-streaming
  ---
  apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
  kind: PubSubTopic
  metadata:
    name: dataflowflextemplatejob-dep1-streaming
  ```


{% endblock %}