# Build
bash build.sh

# Start
docker start <container>

# Stop
docker stop <container>

# Enter container:
docker start <container>
docker exec -it <container> mysql -uroot

# Delete container:
docker rm <container>

# Check log:
docker logs <container>
docker logs <container> 2>&1 | grep GENERATED
