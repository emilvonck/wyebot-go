# wyebot-go

## Usage

Create ".env".
```bash
touch .env
```
Update values below in ".env" to be able to run integration tests.
```bash
WYEBOT_URL=WYEBOT_API_URL # (string) cloud name by looking at the address bar in your browser when logged into the Wyebot dashboard.
WYEBOT_API_KEY=WYEBOT_API_KEY # (string) This can be done from the Management-> API Key. Only dashboard Administrators can generate API Keys.
WYEBOT_LOCATION_ID=1111 # (int) Sample Location ID.
WYEBOT_SENSOR_ID=2222 # (int) Sample Wyebot Sensor ID.
```

Export environment variables.
```bash
export $(cat .env)
```

Run integrations tests.
```bash
go test -v -tags=integration
```