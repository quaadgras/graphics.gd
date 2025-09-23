
function loadJSON() {
    fetch("data.json")
        .then(response => response.json())
        .then(data => {
            // populate #examples div with data
            const examplesDiv = document.getElementById("examples");
            data.forEach(example => {
                const exampleElement = document.createElement("a");
                exampleElement.textContent = example;
                examplesDiv.appendChild(exampleElement);
                exampleElement.href = `${example}`;

                // Add img element
                const imgElement = document.createElement("img");
                imgElement.src = `${example}.webp`;
                exampleElement.appendChild(imgElement);
            });
        })
        .catch(error => console.error("Error loading JSON:", error));
}

document.addEventListener("DOMContentLoaded", function() {
    loadJSON();
});