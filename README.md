# gotzapi
Timezone API written with Go, Echo, and tz.<br><br>
Echo server listens at :31415

# Usage

    $ docker pull docker.pkg.github.com/xlyk/gotzapi/gotzapi:latest
    
    $ docker run -p 31415:31415 --restart unless-stopped -d docker.pkg.github.com/xlyk/gotzapi/gotzapi:latest

# Examples
    
    curl --header "Content-Type: application/json" \
         --request POST \
         --data '{"lat":41.8781,"lon":87.6298}' \
         http://localhost:31415/tz/
