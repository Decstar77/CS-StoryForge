const dropArea = document.getElementById('drop-area');
const dropText = document.getElementById("drop-text");
const fileInput = document.getElementById('file-input');

['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
  dropArea.addEventListener(eventName, preventDefaults, false);
});

function preventDefaults(e) {
  e.preventDefault();
  e.stopPropagation();
}

['dragenter', 'dragover'].forEach(eventName => {
  dropArea.addEventListener(eventName, highlight, false);
});

['dragleave', 'drop'].forEach(eventName => {
  dropArea.addEventListener(eventName, unhighlight, false);
});

function highlight(e) {
  dropArea.classList.add('highlight');
}

function unhighlight(e) {
  dropArea.classList.remove('highlight');
}

dropArea.addEventListener('click', () => {
  fileInput.click();
});

fileInput.addEventListener('change', (e) => {
  loadFile(e.target.files[0]);
});

dropArea.addEventListener('drop', (e) => {
  const files = e.dataTransfer.files;

  if (files[0].name.endsWith('.dem')) {
    loadFile(files[0]);
  } else {
    alert('Please upload a file with a .dem extension.');
  }
}, false);

function readBinaryFileContent(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = function (event) {
      resolve(event.target.result);
    };
    reader.onerror = function (event) {
      reject(event.target.error);
    };
    reader.readAsArrayBuffer(file);
  });
}

function loadFile(file) {
  readBinaryFileContent(file).then(content => {
    dropText.textContent = file.name;
    processFile(content);
  }).catch(error => {
    console.log('Error:', error);
  });
}

function processFile(content) {
  console.log('Processing:', content);
  const demoByteArray = new Uint8Array(content);
  attachGenerateStoryHandler(demoByteArray);
}

function attachGenerateStoryHandler(demoByteArray) {
  const generateButton = document.getElementById('generate-story');
  const outputTextArea = document.getElementById('output-text');
  const worker = new Worker('worker.js');

  worker.addEventListener('message', (e) => {
    const { story } = e.data;
    outputTextArea.value = story;
  });

  generateButton.addEventListener('click', () => {
    outputTextArea.value = "Working...";
    worker.postMessage({ demoByteArray });
  });
}

function generateStory() {
  console.log('Generating story...');

  const storyString = 'Example story: This is a generated story from a Counter Strike demo.';
  return storyString;
}

document.addEventListener('DOMContentLoaded', () => {
  const outputTextArea = document.getElementById('output-text');
  const copyTextButton = document.getElementById('copy-text');

  copyTextButton.addEventListener('click', () => {
    // Copy text to clipboard
    outputTextArea.select();
    document.execCommand('copy');
  
    // Change button text to "Copied!" and deselect text
    const originalButtonText = copyTextButton.textContent;
    copyTextButton.textContent = 'Copied!';
    outputTextArea.setSelectionRange(0, 0);
  
    // Change button text back to original after 1 second
    setTimeout(() => {
      copyTextButton.textContent = originalButtonText;
    }, 3000);
  });
});

