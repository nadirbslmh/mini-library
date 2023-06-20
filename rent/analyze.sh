echo "analyze rent service"

docker run \
--rm \
-e SONAR_HOST_URL="http://localhost:9000" \
-e SONAR_LOGIN="YOUR_TOKEN" \
-v "/path/to/rent:/usr/src" \
--network host \
-m 1g \
sonarsource/sonar-scanner-cli
