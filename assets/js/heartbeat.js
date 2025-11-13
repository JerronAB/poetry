const BACKEND_URL = 'https://log.okpoems.com/heartbeat';
//const BACKEND_URL = 'http://localhost:8080/heartbeat';
const startTime = Date.now();

let visitId;
cookieStore.get('visitId').then(async cookie => {
  if (cookie && cookie.value) {
    // Cookie exists
    console.log(`Cookie Present: ${cookie.value}`)
    visitId = cookie.value;
  } else {
    // Cookie does not exist — create one
    visitId = crypto.randomUUID().substring(0, 6);
    cookieStore.set({
      name: 'visitId',
      value: visitId,
      // optional:
      sameSite: 'lax'
    });
  }
  console.log(`Starting heartbeat for visit: ${visitId}`);
  pulse("init_load")
});

function pulse(state_descriptor) {
    //Calculate time on page in seconds
    const durationOnPage = (Date.now() - startTime) / 1000; // in seconds
    //timezone
    const now = new Date();
    const localTime = now.toLocaleString().replace(", "," ");
    const timeZoneShort = Intl.DateTimeFormat().resolvedOptions().timeZone;
    //Gather page data
    const pageData = {
        visitId: visitId,
        url: window.location.href,
        duration: durationOnPage,
        timestamp: `${localTime} (${timeZoneShort})`,
        referrer: document.referrer,
        stateDesc: state_descriptor
    };
    navigator.sendBeacon(BACKEND_URL, JSON.stringify(pageData));
}

document.addEventListener("visibilitychange", () => {
  if (document.visibilityState === "hidden") {
    console.log(`Sending heartbeat for vis change. `);
    pulse("page_hidden")
  } else if (document.visibilityState === "visible") {
    console.log(`Sending heartbeat for vis change. `);
    pulse("page_viewed")
  }
});