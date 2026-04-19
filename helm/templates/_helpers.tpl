{{/*----------------------------------------------------------
  Expand the name of the chart
------------------------------------------------------------*/}}
{{- define "slurmtalks.chartName" -}}
{{- .Chart.Name -}}
{{- end }}

{{/*----------------------------------------------------------
  Create a default fully qualified app name.
  Truncate at 63 chars because of DNS naming rules.
  If release name contains chart name, use it as full name.
------------------------------------------------------------*/}}
{{- define "slurmtalks.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*----------------------------------------------------------
  Create chart name and version as used by the chart label
------------------------------------------------------------*/}}
{{- define "slurmtalks.chartLabel" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*----------------------------------------------------------
  Common labels
------------------------------------------------------------*/}}
{{- define "slurmtalks.labels" -}}
app.kubernetes.io/name: {{ include "slurmtalks.chartName" . }}
helm.sh/chart: {{ include "slurmtalks.chartLabel" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*----------------------------------------------------------
  Selector labels
------------------------------------------------------------*/}}
{{- define "slurmtalks.selectorLabels" -}}
app.kubernetes.io/name: {{ include "slurmtalks.chartName" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*----------------------------------------------------------
  Image repository helper
  Usage: {{ include "slurmtalks.image" .Values.backend.image }}
------------------------------------------------------------*/}}
{{- define "slurmtalks.image" -}}
{{- printf "%s:%s" .repository (.tag | default "latest") -}}
{{- end -}}

{{/*----------------------------------------------------------
  Fullnames for sub chars
------------------------------------------------------------*/}}
{{- define "slurmtalks.frontendFullname" -}}
{{- printf "%s-frontend" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- end }}

{{- define "slurmtalks.backendFullname" -}}
{{- printf "%s-backend" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- end }}

