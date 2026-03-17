//these are here so they stay in the global scope
let pyodide;

async function loadAndRun(poems_content){
    stat = document.getElementById('wordcounter')
    pyodide = await loadPyodide();
    try {
        const resp = await fetch('/CountAndHighlight.py');
        const code = await resp.text();
        pyodide.globals.set("poems_content", poems_content);
        let result = await pyodide.runPython(code);
        console.log(`Result: ${result}`)
        // redeclaring this here bc I'm not sure
        // how global variables are scoped in js
        let paragraphs = Array.from(document.querySelectorAll('main p')).slice(1);
        for (const word of result) {
            for (let p of paragraphs) {
                p.innerHTML = p.innerHTML.replaceAll(word,`<h-me>${word}</h-me>`)
                let proper_version = word.charAt(0).toUpperCase() + word.slice(1)
                p.innerHTML = p.innerHTML.replaceAll(proper_version,`<h-me>${proper_version}</h-me>`)
            }
        }
    } catch (err) {
        console.error(err);
    }
}
// collects all paragraph text inside "main" except the first one
const paragraphs = Array.from(document.querySelectorAll('main p')).slice(1);
const poems_content = paragraphs.map(p => p.textContent.trim()).join("\n\n");
loadAndRun(poems_content);
