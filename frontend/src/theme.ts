import type { GlobalThemeOverrides } from 'naive-ui'

export const themeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: '#A8E6CF',
    primaryColorHover: '#88D4AB',
    primaryColorPressed: '#6BC49B',
    fontFamily: "'PingFang SC', 'Microsoft YaHei', 'HarmonyOS Sans', system-ui, sans-serif",
    borderRadius: '12px',
  },
  Button: {
    borderRadiusMedium: '12px',
    borderRadiusLarge: '16px',
  },
  Input: {
    borderRadius: '12px',
  },
  Card: {
    borderRadius: '16px',
  },
}
