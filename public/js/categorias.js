const apiUrl = "http://localhost:8080";

document.getElementById("createCategoriaForm").addEventListener("submit", async (e) => {
  e.preventDefault();
  const nome = document.getElementById("nome").value;

  await fetch(`${apiUrl}/categorias/create`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ nome })
  });

  alert("Categoria criada com sucesso!");
  loadCategorias();
});

async function loadCategorias() {
  const res = await fetch(`${apiUrl}/categorias/read`);
  const categorias = await res.json();
  const tbody = document.getElementById("categoriasTable");
  tbody.innerHTML = "";
  categorias.forEach(c => {
    const tr = document.createElement("tr");
    tr.innerHTML = `
      <td>${c.id}</td>
      <td>${c.nome}</td>
      <td>
        <button onclick="deleteCategoria(${c.id})">Excluir</button>
      </td>
    `;
    tbody.appendChild(tr);
  });
}

async function deleteCategoria(id) {
  await fetch(`${apiUrl}/categorias/delete?id=${id}`, { method: "DELETE" });
  alert("Categoria deletada!");
  loadCategorias();
}

loadCategorias();
