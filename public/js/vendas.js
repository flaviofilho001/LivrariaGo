const apiUrl = "http://localhost:8080";

document.getElementById("createVendaForm").addEventListener("submit", async (e) => {
  e.preventDefault();
  const data = {
    cliente_id: parseInt(document.getElementById("cliente_id").value),
    valor_total: document.getElementById("valor_total").value,
    forma_pagamento: document.getElementById("forma_pagamento").value,
    status: document.getElementById("status").value,
    observacoes: document.getElementById("observacoes").value
  };

  await fetch(`${apiUrl}/vendas/create`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data)
  });

  alert("Venda criada com sucesso!");
  loadVendas();
});

async function loadVendas() {
  const res = await fetch(`${apiUrl}/vendas/read`);
  const vendas = await res.json();
  const tbody = document.getElementById("vendasTable");
  tbody.innerHTML = "";
  vendas.forEach(v => {
    const tr = document.createElement("tr");
    tr.innerHTML = `
      <td>${v.id}</td>
      <td>${v.cliente_id}</td>
      <td>${v.valor_total}</td>
      <td>${v.forma_pagamento}</td>
      <td>${v.status}</td>
      <td>${v.observacoes}</td>
      <td><button onclick="deleteVenda(${v.id})">Excluir</button></td>
    `;
    tbody.appendChild(tr);
  });
}

async function deleteVenda(id) {
  await fetch(`${apiUrl}/vendas/delete?id=${id}`, { method: "DELETE" });
  alert("Venda deletada!");
  loadVendas();
}

loadVendas();
