function startHeartbeatTracking() {
    //URL for the backend endpoint
    const BACKEND_URL = 'https://log.okpoems.com/heartbeat';
    //const BACKEND_URL = 'http://localhost:8080/heartbeat';
    const visitId = crypto.randomUUID();
    const startTime = Date.now();

    console.log(`Starting heartbeat tracking for visit: ${visitId}`);

    async function sendHeartbeat() {
        //Calculate time on page in seconds
        const durationOnPage = (Date.now() - startTime) / 1000; // in seconds
        
        //Gather page data
        const pageData = {
            visitId: visitId,
            url: window.location.href,
            duration: durationOnPage,
            timestamp: new Date().toISOString(),
            referrer: document.referrer
        };

        try {
            const response = await fetch(BACKEND_URL, {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify(pageData)
            });

            if (!response.ok) {
                console.error('Heartbeat request failed:', response.statusText);
            }
        } catch (error) {
            console.error('Error sending heartbeat:', error);
        }
    }

    //Send the first heartbeat immediately and then every 5 seconds
    sendHeartbeat();
    setInterval(sendHeartbeat, 5000);
}

startHeartbeatTracking()
