+++
date = '2025-04-21T09:19:46-05:00'
draft = false
title = ''
framed = true
+++
This page is a list of all the poems you can find on this site, sorted by date. The most recent is at the top, and the oldest is at the bottom. 

{{< rawhtml >}}
<style>
  body {
    transition: background-color 0.1s linear, color 0.3s ease;
  }
  h2 {
    transition: color 0.3s ease;
  }
</style>

<script>
  maxScrollPercent = 0;
  window.addEventListener('scroll', () => {
    const scrollTop = window.scrollY;
    const docHeight = document.documentElement.scrollHeight - window.innerHeight;
    const scrollPercent = Math.min(scrollTop / docHeight, 1);
    
    if (scrollPercent > maxScrollPercent) {
        // interpolate between white (255) and dark blue (0,31,63)
        const r = Math.round(255 - scrollPercent * (255 - 0));
         const g = Math.round(255 - scrollPercent * (255 - 31));
         const b = Math.round(255 - scrollPercent * (255 - 63));
        document.body.style.backgroundColor = `rgb(${r}, ${g}, ${b})`;

        // make text gradually lighten
        const textColor = Math.round(0 + scrollPercent * (255 - 0));
        const textRGB = `rgb(${textColor}, ${textColor}, ${textColor})`;

        document.body.style.color = textRGB;

        // update all h2 colors
        document.querySelectorAll('a').forEach(a => {
        a.style.color = textRGB;
        });
        maxScrollPercent = scrollPercent;
    }
  });
</script>
{{</ rawhtml >}}