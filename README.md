
# cloudflare-liwanetwork

Application which can update the cloudflare DNS records from the public IP of the network gateway. Can be used mostly in local or home network setup when ISP doesn't entertain your static public request.

This program monitor for your ip changes and any changes will be reflected in cloudflare records 

    <html>
	    <head>
	    </head>
    </html>


## Running

*cloudflare zone name*

    export ZONE_NAME=example.com

*api token*

    export API_TOKEN=***********

 

records that needs to be updated as can be passed as arguments to program

    ./liwa-nw example.com link.example.com
