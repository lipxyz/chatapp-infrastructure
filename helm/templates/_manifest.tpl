{{- define "manifest.template" -}}
apiVersion: {{ .config.apiVersion }}
kind: {{ .config.kind }}
metadata:
  name: {{ include "slurmtalks.fullname" .root }}
  labels:
    {{- include "slurmtalks.labels" .root | nindent 4 }}

{{- if eq .config.kind "Deployment" }}
{{- $enabled := true }}
{{- if hasKey .config "enabled" }}
  {{- $enabled = .config.enabled }}
{{- end }}
{{- if $enabled }}
spec:
  replicas: {{ if hasKey .config "replicas" }}{{ .config.replicas }}{{ else }}1{{ end }}
  strategy:
    type: {{ .config.strategy.type | default "RollingUpdate" }}
    {{- if eq (.config.strategy.type | default "RollingUpdate") "RollingUpdate" }}
    rollingUpdate:
      maxSurge: {{ .config.strategy.rollingUpdate.maxSurge | default 1 }}
      maxUnavailable: {{ .config.strategy.rollingUpdate.maxUnavailable | default 0 }}
    {{- end }}
  selector:
    matchLabels:
      {{- include "slurmtalks.selectorLabels" .root | nindent 6 }}
  template:
    metadata:
      labels:
        {{- include "slurmtalks.selectorLabels" .root | nindent 8 }}
      annotations:
        {{- with .config.prometheus }}
        {{- if .scrape }}
        prometheus.io/scrape: "true"
        {{- if .path }}
        prometheus.io/path: "{{ .path }}"
        {{- end }}
        {{- if .port }}
        prometheus.io/port: "{{ .port }}"
        {{- end }}
        {{- end }}
        {{- end }}
    spec:
      {{- with .config.serviceAccountName }}
      serviceAccountName: {{ . }}
      {{- end }}
      {{- with .config.imagePullSecrets }}
      imagePullSecrets:
        {{- range . }}
        - name: {{ .name }}
        {{- end }}
      {{- end }}
      {{- if .config.nodeSelector }}
      nodeSelector:
        {{- toYaml .config.nodeSelector | nindent 8 }}
      {{- end }}
      {{- if .config.affinity }}
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
              - matchExpressions:
                  {{- range .config.affinity.matchExpressions }}
                  - key: {{ .key }}
                    operator: {{ .operator }}
                    values:
                      {{- range .values }}
                      - {{ . }}
                      {{- end }}
                  {{- end }}
      {{- end }}
      containers:
        - name: {{ .config.name }}
          image: {{ include "slurmtalks.image" .config.image }}
          imagePullPolicy: {{ .config.image.pullPolicy | default "IfNotPresent" }}
          {{- if .config.ports }}
          ports:
            {{- range .config.ports }}
            - containerPort: {{ .containerPort }}
            {{- end }}
          {{- end }}
          {{- if .config.env }}
          env:
            {{- range .config.env }}
            - name: {{ .name }}
              {{- if .value }}
              value: {{ .value | quote }}
              {{- end }}
              {{- if .valueFrom }}
              valueFrom:
                {{- toYaml .valueFrom | nindent 16 }}
              {{- end }}
            {{- end }}
          {{- end }}
          {{- if .config.volumeMounts }}
          volumeMounts:
            {{- range .config.volumeMounts }}
            - name: {{ .name }}
              mountPath: {{ .mountPath }}
            {{- end }}
          {{- end }}
          {{- if .config.resources }}
          resources:
            {{- toYaml .config.resources | nindent 12 }}
          {{- end }}
          {{- if .config.livenessProbe }}
          livenessProbe:
            {{- toYaml .config.livenessProbe | nindent 12 }}
          {{- end }}
          {{- if .config.readinessProbe }}
          readinessProbe:
            {{- toYaml .config.readinessProbe | nindent 12 }}
          {{- end }}
      {{- if .config.volumes }}
      volumes:
        {{- range .config.volumes }}
        - name: {{ .name }}
          {{- if .persistentVolumeClaim }}
          persistentVolumeClaim:
            claimName: {{ .persistentVolumeClaim.claimName }}
          {{- else if .emptyDir }}
          emptyDir: {}
          {{- end }}
        {{- end }}
      {{- end }}
{{- end }}

{{- else if eq .config.kind "Service" }}
{{- $enabled := true }}
{{- if hasKey .config "enabled" }}
  {{- $enabled = .config.enabled }}
{{- end }}
{{- if $enabled }}
spec:
  selector:
    {{- include "slurmtalks.selectorLabels" .root | nindent 4 }}
  {{- if .config.ports }}
  ports:
    {{- range .config.ports }}
    - port: {{ .port }}
      targetPort: {{ .targetPort }}
    {{- end }}
  {{- end }}
  type: {{ .config.type | default "ClusterIP" }}
{{- end }}
{{- end }}
{{- end }}