function switchTab(tabName) {
    // Remove active class from all tabs and tab contents
    document.querySelectorAll('.tab').forEach(tab => tab.classList.remove('active'));
    document.querySelectorAll('.tab-content').forEach(content => content.classList.remove('active'));
    
    // Add active class to clicked tab and corresponding content
    event.target.classList.add('active');
    document.getElementById(tabName + 'Tab').classList.add('active');
    
    // Hide results when switching tabs
    document.getElementById('lettersResultsContainer').style.display = 'none';
    document.getElementById('numbersResultsContainer').style.display = 'none';
    
    // Update button states
    updateButtonStates();
}

// Letters Game Functions
function handleLetterInput(input, index) {
    // Only allow letters
    const value = input.value.replace(/[^a-zA-Z]/g, '').toUpperCase();
    input.value = value;
    
    // Auto-focus next input if current has value and not last input
    if (value && index < 8) {
        const nextInput = input.parentElement.children[index + 1];
        if (nextInput) {
            nextInput.focus();
        }
    }
    
    updateButtonStates();
}

function resetLettersGame() {
    const inputs = document.querySelectorAll('.letter-box');
    inputs.forEach(input => input.value = '');
    
    document.getElementById('lettersResultsContainer').style.display = 'none';
    
    if (inputs.length > 0) {
        inputs[0].focus();
    }
    
    updateButtonStates();
}

function getLetters() {
    const inputs = document.querySelectorAll('.letter-box');
    return Array.from(inputs)
        .map(input => input.value.trim().toLowerCase())
        .filter(letter => letter !== '');
}

async function findWords() {
    const letters = getLetters();
    
    if (letters.length === 0) {
        alert('Please enter at least one letter!');
        return;
    }

    const resultsContainer = document.getElementById('lettersResultsContainer');
    const wordsGrid = document.getElementById('wordsGrid');
    const button = document.getElementById('findWordsBtn');
    
    button.disabled = true;
    button.textContent = 'Finding Words...';
    resultsContainer.style.display = 'block';
    wordsGrid.innerHTML = '<div class="loading">🔍 Searching for words...</div>';

    try {
        const lettersParam = letters.join(';');
        const response = await fetch(`/words/${lettersParam}`);
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const data = await response.json();
        displayWordsResults(data);
        
    } catch (error) {
        console.error('Error fetching words:', error);
        wordsGrid.innerHTML = `
            <div class="error">
                <strong>Error:</strong> Could not connect to the API. 
                Make sure the server is running on localhost:3000
                <br><br>
                <small>Error details: ${error.message}</small>
            </div>
        `;
    } finally {
        button.disabled = false;
        button.textContent = 'Find Words';
    }
}

function displayWordsResults(data) {
    const wordsGrid = document.getElementById('wordsGrid');
    
    if (!data.dictionary || data.dictionary.length === 0) {
        wordsGrid.innerHTML = `
            <div class="word-card">
                <div class="word">No words found</div>
                <div class="definition">Try different letters or check your spelling!</div>
            </div>
        `;
        return;
    }

    let html = '';
    data.dictionary.forEach((word, index) => {
        const definition = data.definitions[index] || 'No definition available';
        html += `
            <div class="word-card">
                <div class="word">${word}</div>
                <div class="definition">${definition}</div>
            </div>
        `;
    });
    
    wordsGrid.innerHTML = html;
}

// Numbers Game Functions
function handleNumberInput(input, index) {
    // Only allow numbers, max 2 characters
    const value = input.value.replace(/[^0-9]/g, '').slice(0, 2);
    input.value = value;
    
    // Do not auto-focus - user must press TAB to move to next input
    updateButtonStates();
}

function handleTargetInput(input) {
    // Only allow numbers
    const value = input.value.replace(/[^0-9]/g, '');
    input.value = value;
    updateButtonStates();
}

function resetNumbersGame() {
    const numberInputs = document.querySelectorAll('.number-box');
    const targetInput = document.getElementById('targetBox');
    
    numberInputs.forEach(input => input.value = '');
    targetInput.value = '';
    
    document.getElementById('numbersResultsContainer').style.display = 'none';
    
    if (numberInputs.length > 0) {
        numberInputs[0].focus();
    }
    
    updateButtonStates();
}

function getNumbers() {
    const inputs = document.querySelectorAll('.number-box');
    return Array.from(inputs)
        .map(input => input.value.trim())
        .filter(num => num !== '')
        .map(num => parseInt(num));
}

function getTarget() {
    const target = document.getElementById('targetBox').value.trim();
    return target ? parseInt(target) : null;
}

async function solveNumbers() {
    const numbers = getNumbers();
    const target = getTarget();
    
    if (numbers.length !== 6) {
        alert('Please enter exactly 6 numbers!');
        return;
    }
    
    if (!target) {
        alert('Please enter a target number!');
        return;
    }

    const resultsContainer = document.getElementById('numbersResultsContainer');
    const solutionsGrid = document.getElementById('solutionsGrid');
    const button = document.getElementById('solveNumbersBtn');
    
    button.disabled = true;
    button.textContent = 'Solving...';
    resultsContainer.style.display = 'block';
    solutionsGrid.innerHTML = '<div class="loading">🧮 Calculating solutions...</div>';

    try {
        const numbersParam = numbers.join(',');
        const response = await fetch(`/numbers/${numbersParam}/${target}`);
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const data = await response.json();
        displayNumbersResults(data);
        
    } catch (error) {
        console.error('Error solving numbers:', error);
        solutionsGrid.innerHTML = `
            <div class="error">
                <strong>Error:</strong> Could not connect to the API. 
                Make sure the server is running on localhost:3000
                <br><br>
                <small>Error details: ${error.message}</small>
            </div>
        `;
    } finally {
        button.disabled = false;
        button.textContent = 'Solve Numbers';
    }
}

function displayNumbersResults(data) {
    const solutionsGrid = document.getElementById('solutionsGrid');
    
    if (!data.solutions || data.solutions.length === 0) {
        solutionsGrid.innerHTML = `
            <div class="solution-card">
                <div class="expression">No solutions found</div>
                <div class="result">Try different numbers or a different target!</div>
            </div>
        `;
        return;
    }

    let html = '';
    data.solutions.forEach((solution, index) => {
        const isExact = solution.distance === 0;
        const cardClass = isExact ? 'solution-card exact-match' : 'solution-card';
        
        html += `
            <div class="${cardClass}">
                <div class="expression">${solution.expression}</div>
                <div class="result">Result: ${solution.result}</div>
                <div class="result">Distance from target: ${solution.distance}</div>
                ${isExact ? '<div class="result" style="color: #00ff88;">🎯 Exact Match!</div>' : ''}
            </div>
        `;
    });
    
    solutionsGrid.innerHTML = html;
}

function updateButtonStates() {
    // Letters game button
    const hasLetters = getLetters().length > 0;
    const findWordsBtn = document.getElementById('findWordsBtn');
    if (findWordsBtn) findWordsBtn.disabled = !hasLetters;
    
    // Numbers game button
    const numbers = getNumbers();
    const target = getTarget();
    const hasValidNumbers = numbers.length === 6 && target !== null;
    const solveNumbersBtn = document.getElementById('solveNumbersBtn');
    if (solveNumbersBtn) solveNumbersBtn.disabled = !hasValidNumbers;
}

// Add keydown event to prevent non-letter/number input
document.addEventListener('keydown', function(event) {
    if (event.target.classList.contains('letter-box')) {
        // Allow backspace, delete, tab, escape, enter
        if ([8, 9, 27, 13, 46].indexOf(event.keyCode) !== -1 ||
            // Allow Ctrl+A, Ctrl+C, Ctrl+V, Ctrl+X
            (event.keyCode === 65 && event.ctrlKey === true) ||
            (event.keyCode === 67 && event.ctrlKey === true) ||
            (event.keyCode === 86 && event.ctrlKey === true) ||
            (event.keyCode === 88 && event.ctrlKey === true)) {
            return;
        }
        // Ensure that it is a letter
        if ((event.shiftKey || (event.keyCode < 65 || event.keyCode > 90))) {
            event.preventDefault();
        }
    }
    
    if (event.target.classList.contains('number-box') || event.target.classList.contains('target-box')) {
        // Allow backspace, delete, tab, escape, enter
        if ([8, 9, 27, 13, 46].indexOf(event.keyCode) !== -1 ||
            // Allow Ctrl+A, Ctrl+C, Ctrl+V, Ctrl+X
            (event.keyCode === 65 && event.ctrlKey === true) ||
            (event.keyCode === 67 && event.ctrlKey === true) ||
            (event.keyCode === 86 && event.ctrlKey === true) ||
            (event.keyCode === 88 && event.ctrlKey === true)) {
            return;
        }
        // Ensure that it is a number
        if (event.keyCode < 48 || event.keyCode > 57) {
            event.preventDefault();
        }
    }
    
    // Handle Enter key for games
    if (event.key === 'Enter') {
        const activeTab = document.querySelector('.tab-content.active');
        if (activeTab.id === 'lettersTab') {
            const button = document.getElementById('findWordsBtn');
            if (button && !button.disabled) {
                findWords();
            }
        } else if (activeTab.id === 'numbersTab') {
            const button = document.getElementById('solveNumbersBtn');
            if (button && !button.disabled) {
                solveNumbers();
            }
        }
    }
});

// Initialize button states when page loads
document.addEventListener('DOMContentLoaded', function() {
    updateButtonStates();
});