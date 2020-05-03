# gotzapi
Timezone API written with Go, Echo, and tz.<br><br>
Echo server listens at :31415

# Usage

    $ ./gotzapi

# Examples
    
    curl --header "Content-Type: application/json" \
         --request POST \
         --data '{"lat":41.8781,"lon":87.6298}' \
         http://localhost:31415/tz/
