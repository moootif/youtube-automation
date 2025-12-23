// This file contains JavaScript code to control the dynamic behavior of the HTML page.

document.addEventListener("DOMContentLoaded", function () {
  const btn = document.getElementById("runBtn");
  const result = document.getElementById("result");

  btn.addEventListener("click", async () => {
    btn.disabled = true;
    result.textContent = "実行中...";
    try {
      const res = await fetch("/api/run", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ action: "upload" })
      });
      if (!res.ok) throw new Error("サーバーエラー");
      const data = await res.json();
      result.textContent = data.result || "完了";
    } catch (e) {
      result.textContent = "エラー: " + e.message;
    } finally {
      btn.disabled = false;
    }
  });
});