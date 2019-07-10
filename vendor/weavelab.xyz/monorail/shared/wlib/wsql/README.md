# WSQL

## Cloud SQL Proxy

To enable the use of Cloud SQL Proxy, the following conditions need to be met:

1) If connecting to an instance with ONLY a Private IP, the app must be running in the same VPC to connect
2) The app must have Root SSL certs mounted:

```yaml
spec:
  template:
    spec:
      containers:
      - name: app
        volumeMounts:
        - mountPath: /etc/ssl/certs
          name: sslroots
      volumes:
      - hostPath:
          path: /etc/ssl/certs
        name: sslroots
```

3) If running in a container, must have valid credentials for a service account available and the `GOOGLE_APPLICATION_CREDENTIALS` env set:

```yaml
spec:
  template:
    spec:
      containers:
      - name: app
        env:
        - name: GOOGLE_APPLICATION_CREDENTIALS
          value: /var/secrets/google/key.json
        volumeMounts:
        - mountPath: /var/secrets/google
          name: google-cloud-key
      volumes:
      - name: google-cloud-key
        secret:
          defaultMode: 420
          optional: false
          secretName: cloud-sql-proxy-sa
```

4) Hostname of the connection string must be of the form `project:region:instance`.
5) SSL mode on the connection must be disabled. The proxy handles SSL for you.

To enable your application to use the Cloud SQL Proxy, change the `Driver` param on the `ConnectionString` struct:

```go
DBSettings := &wsql.Settings{
    PrimaryConnectString: wsql.ConnectString{
        Driver: wsql.CloudSQLDriver,
    },
}
```