disable_snapshots()
analytics_settings(enable=False)
if os.environ.get('TILT_ALLOWED_CONTEXT', "") != "":
    allow_k8s_contexts(os.environ.get('TILT_ALLOWED_CONTEXT', ""))

k8s_yaml('kubernetes.yaml')
k8s_yaml('minio.yaml')

k8s_resource(
  'nginx',
  port_forwards=['8080:80'],
)

k8s_resource(
  'minio',
  port_forwards=['9000:9000'],
)
