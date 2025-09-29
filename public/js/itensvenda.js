const apiUrl = "http://localhost:8080";

document.getElementById("createItemVendaForm").addEventListener("submit", async (e) => {
  e.preventDefault();
  const quantidade = parseInt(document.getElementById("quantidade").value);
  const preco_unitario = parseFloat(document.getElementById("preco_unitario").value);
  const desconto = parseFloat(document.getElementById("desconto").value || 0);

  const data = {
    venda_id: parseInt(document.getElementById("venda_id").value),
    livro_id: parseInt(document.getElementById("livro_id").value),
    quantidade,
    preco_unitario,
    desconto
  };

  await fetch(`${apiUrl}/itensvenda/create`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data)
  });

  alert("Item de Venda criado com sucesso!");
  loadItensVenda();
});

async function loadItensVenda() {
  const res = await fetch(`${apiUrl}/itensvenda/read`);
  const itens = await res.json();
  const tbody = document.getElementById("itensVendaTable");
  tbody.innerHTML = "";
  itens.forEach(i => {
    const tr = document.createElement("tr");
    tr.innerHTML = `
      <td>${i.id}</td>
      <td>${i.venda_id}</td>
      <td>${i.livro_id}</td>
      <td>${i.quantidade}</td>
      <td>${i.preco_unitario}</td>
      <td>${i.subtotal}</td>
      <td>${i.desconto}</td>
      <td><button onclick="deleteItemVenda(${i.id})">Excluir</button></td>
    `;
    tbody.appendChild(tr);
  });
}

async function deleteItemVenda(id) {
  await fetch(`${apiUrl}/itensvenda/delete?id=${id}`, { method: "DELETE" });
  alert("Item de Venda deletado!");
  loadItensVenda();
}

loadItensVenda();
