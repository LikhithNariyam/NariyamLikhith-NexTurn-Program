const expenseForm = document.getElementById("expense-form");
const expenseTableBody = document.querySelector("#expense-table tbody");
const expenseChart = document.getElementById("expense-chart");

let expenses = JSON.parse(localStorage.getItem("expenses")) || [];
let chart = null;

function updateUI() {
  // Clear the table
  expenseTableBody.innerHTML = "";

  // Populate the table
  expenses.forEach((expense, index) => {
    const row = document.createElement("tr");
    row.innerHTML = `
      <td>${expense.amount}</td>
      <td>${expense.description}</td>
      <td>${expense.category}</td>
      <td>
        <button onclick="deleteExpense(${index})">Delete</button>
      </td>
    `;
    expenseTableBody.appendChild(row);
  });

  // Update chart
  updateChart();

  // Save to localStorage
  localStorage.setItem("expenses", JSON.stringify(expenses));
}

// Add a new expense
expenseForm.addEventListener("submit", (e) => {
  e.preventDefault();

  const amount = parseFloat(document.getElementById("amount").value);
  const description = document.getElementById("description").value.trim();
  const category = document.getElementById("category").value;

  if (!isNaN(amount) && description && category) {
    expenses.push({ amount, description, category });
    expenseForm.reset();
    updateUI();
  }
});

// Delete an expense
function deleteExpense(index) {
  expenses.splice(index, 1);
  updateUI();
}

function updateChart() {
  const categoryTotals = expenses.reduce((totals, expense) => {
    totals[expense.category] = (totals[expense.category] || 0) + expense.amount;
    return totals;
  }, {});

  const labels = Object.keys(categoryTotals);
  const data = Object.values(categoryTotals);

  if (chart) chart.destroy();

  chart = new Chart(expenseChart, {
    type: "pie", // You can change to "bar" if preferred
    data: {
      labels: labels,
      datasets: [
        {
          label: "Expenses",
          data: data,
          backgroundColor: ["#007bff", "#28a745", "#ffc107", "#dc3545", "#6c757d"],
        },
      ],
    },
    options: {
      responsive: true,
    },
  });
}

updateUI();
