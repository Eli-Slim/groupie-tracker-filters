document.getElementById("search-input").addEventListener("input", function() {
    const query = this.value;
    const suggestionsDiv = document.getElementById("suggestions");

    if (query.length > 0) {
        fetch(`/suggestions?search=${encodeURIComponent(query)}`)
            .then(response => response.json())
            .then(data => {
                suggestionsDiv.innerHTML = "";
                if (data.length > 0) {
                    data.forEach(suggestion => {
                        for (const [key, value] of Object.entries(suggestion)) {
                            const suggestionItem = document.createElement("div");
                            suggestionItem.classList.add("suggestion-item");
                            suggestionItem.textContent = key;
                            suggestionItem.addEventListener("click", () => {
                                window.location.href = value;
                            });
                            suggestionsDiv.appendChild(suggestionItem);    
                        }
                    });
                } else {
                    suggestionsDiv.innerHTML = "";
                }
            })
            .catch(error => {
                console.error('Error fetching suggestions:', error);
            });
    } else {
        suggestionsDiv.innerHTML = "";
    }
});
