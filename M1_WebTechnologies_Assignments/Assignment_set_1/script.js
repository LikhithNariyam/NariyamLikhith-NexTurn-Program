const taskInput = document.getElementById("task-input");
const addTaskBtn = document.getElementById("add-task-btn");
const taskList = document.getElementById("task-list");
const pendingCountEl = document.getElementById("pending-count");

// Load tasks from localStorage
let tasks = JSON.parse(localStorage.getItem("tasks")) || [];

// Update the UI with tasks
function updateUI() {
  taskList.innerHTML = "";
  let pendingCount = 0;

  tasks.forEach((task, index) => {
    const li = document.createElement("li");
    li.className = `task-item ${task.completed ? "completed" : ""}`;
    li.draggable = true; // Enable drag-and-drop
    li.dataset.index = index; // Store the task index for drag-and-drop

    li.innerHTML = `
      <span>${task.name}</span>
      <div>
        <button onclick="editTask(${index})">Edit</button>
        <button onclick="deleteTask(${index})">Delete</button>
        <button onclick="toggleComplete(${index})">✔</button>
      </div>
    `;

    // Add drag-and-drop event listeners
    li.addEventListener("dragstart", handleDragStart);
    li.addEventListener("dragover", handleDragOver);
    li.addEventListener("drop", handleDrop);
    li.addEventListener("dragend", handleDragEnd);

    taskList.appendChild(li);

    if (!task.completed) pendingCount++;
  });

  pendingCountEl.textContent = `Pending Tasks: ${pendingCount}`;
  localStorage.setItem("tasks", JSON.stringify(tasks));
}

// Add a new task
addTaskBtn.addEventListener("click", () => {
  const taskName = taskInput.value.trim();
  if (taskName) {
    tasks.push({ name: taskName, completed: false });
    taskInput.value = "";
    updateUI();
  }
});

// Edit a task
function editTask(index) {
  const newTaskName = prompt("Edit your task:", tasks[index].name);
  if (newTaskName) {
    tasks[index].name = newTaskName.trim();
    updateUI();
  }
}

// Delete a task
function deleteTask(index) {
  tasks.splice(index, 1);
  updateUI();
}

// Toggle task completion
function toggleComplete(index) {
  tasks[index].completed = !tasks[index].completed;
  updateUI();
}

// Drag-and-Drop Handlers
let draggedIndex = null;

function handleDragStart(event) {
  draggedIndex = +event.target.dataset.index; // Get the index of the dragged task
  event.target.style.opacity = "0.5"; // Make the dragged item semi-transparent
}

function handleDragOver(event) {
  event.preventDefault(); // Allow dropping
  event.target.style.border = "2px dashed #007bff"; // Add visual feedback
}

function handleDrop(event) {
  event.preventDefault();
  const droppedIndex = +event.target.dataset.index; // Get the index of the dropped-on task

  if (draggedIndex !== null && droppedIndex !== undefined) {
    // Swap the tasks in the array
    const draggedTask = tasks[draggedIndex];
    tasks.splice(draggedIndex, 1);
    tasks.splice(droppedIndex, 0, draggedTask);
    updateUI(); // Refresh the UI
  }
}

function handleDragEnd(event) {
  event.target.style.opacity = "1"; // Reset opacity
  Array.from(taskList.children).forEach((child) => {
    child.style.border = "none"; // Reset borders
  });
}

// Initialize UI
updateUI();
