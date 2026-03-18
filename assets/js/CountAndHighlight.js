//these are here so they stay in the global scope
const go = new Go();

// collects all paragraph text inside "main" except the first one
const paragraphs_ = Array.from(document.querySelectorAll('main p')).slice(1);
const poems_content = paragraphs_.map(p => p.textContent.trim()).join("\n\n");

// Fetch and instantiate the WASM module
WebAssembly.instantiateStreaming(fetch("/wasm.wasm"), go.importObject)
    .then((result) => {
        // Start the Go program (it will block and listen)
        go.run(result.instance);
        // Call the function exported from Go
        const words = window.getEligibleWords(poems_content);
        console.log(words)
        // Render the returned JavaScript array
        let paragraphs = Array.from(document.querySelectorAll('main p')).slice(1);
        for (const word of words) {
            for (let p of paragraphs) {
                p.innerHTML = p.innerHTML.replaceAll(word,`<h-me>${word}</h-me>`)
                let proper_version = word.charAt(0).toUpperCase() + word.slice(1)
                p.innerHTML = p.innerHTML.replaceAll(proper_version,`<h-me>${proper_version}</h-me>`)
            }
        }
    });
