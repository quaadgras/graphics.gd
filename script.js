
function loadJSON() {
    fetch("data.json")
        .then(response => response.json())
        .then(data => {
            // populate #examples div with data
            const examplesDiv = document.getElementById("examples");
            data.forEach(example => {
                const exampleElement = document.createElement("div");
                exampleElement.textContent = example;
                examplesDiv.appendChild(exampleElement);
            });
        })
        .catch(error => console.error("Error loading JSON:", error));
}

document.addEventListener("DOMContentLoaded", function() {
    loadJSON();
});