const output = document.querySelector("#output")
const shortenButton = document.querySelector("#shortenButton")

console.log("URL SHORTENER")

const templateString = "Your Shortened URL is: "

shortenButton.onclick = async function () {
    const urlInput = document.querySelector("#urlInput")
    const shortUrlInput = document.querySelector("#shortUrlInput")
    const url = urlInput.value
    const shortUrl = shortUrlInput.value
    const res = await fetch(`http://localhost:${PORT}/create`, {
        method: 'post',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            url,
            shortUrl
        })
    })
    const json = await res.json()
    if (json.error !== undefined){
        output.innerHTML = json.error
    } else {
        console.log(json)
        const shortUrlOutput = json['short_url']
        output.innerHTML = `${templateString}<a href="http://${HOSTNAME}/${shortUrlOutput}" target="_blank">http://${HOSTNAME}/${shortUrlOutput}</a>`   
    }
}