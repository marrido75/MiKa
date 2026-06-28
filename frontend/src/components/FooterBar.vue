<template>
  <footer class="footer">
    <div class="footer-inner">
      <div class="footer-brand">✦ MiKa</div>
      <div class="footer-links">
        <a href="#" class="footer-link">帮助中心</a>
        <a href="#" class="footer-link">联系我们</a>
        <a href="#" class="footer-link">服务条款</a>
      </div>
      <div class="footer-copy">© 2026 MiKa. All rights reserved.</div>
    </div>
  </footer>
</template>

<style scoped>
.footer {
  border-top: 1px solid var(--color-border);
  background-color: rgba(255, 255, 255, 0.6);
}

.footer-inner {
  max-width: 1200px;
  margin: 0 auto;
  padding: 32px 24px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
}

.footer-brand {
  font-size: 18px;
  font-weight: 700;
  color: var(--color-text);
  font-family: var(--font-serif);
}

.footer-links {
  display: flex;
  gap: 24px;
}

.footer-link {
  font-size: 13px;
  color: var(--color-text-secondary);
  text-decoration: none;
  transition: color 0.2s ease;
}

.footer-link:hover {
  color: var(--color-primary-pressed);
}

.footer-copy {
  font-size: 12px;
  color: var(--color-text-secondary);
  opacity: 0.7;
}
</style>
