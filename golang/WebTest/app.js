const baseURL = 'https://4950-124-120-5-94.ngrok.io/'
fetch(baseURL)
.then(resp => resp.json())
.then(data => displayData(data));

function displayData(data) {
    document.querySelector("pre").innerHTML = JSON.stringify(data, null, 2);
}