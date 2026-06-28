# MiKa 发卡网站设计文档

## 项目概述

MiKa 是一个简约清新的虚拟卡/密钥销售平台，用户可以浏览商品、购买卡密、管理订单。

## 设计风格

**方案 A：清新渐变风格**

### 配色系统

| 角色 | 色值 | 用途 |
|------|------|------|
| 主色 | `#A8E6CF` | 按钮、链接、重点元素 |
| 主色深 | `#88D4AB` | hover 状态、渐变 |
| 强调色 | `#FFB7B2` | 价格、徽章、CTA |
| 背景色 | `#FFFBF5` | 页面背景 |
| 卡片背景 | `#FFFFFF` | 商品卡片 |
| 文字主色 | `#2D3436` | 标题、正文 |
| 文字次色 | `#636E72` | 描述、辅助信息 |
| 边框色 | `#DFE6E9` | 分割线、边框 |

### 字体

- 标题：Noto Serif SC（衬线，增添文艺感）
- 正文：Noto Sans SC（无衬线，清晰易读）
- 代码/数据：monospace

### 签名元素

商品卡片使用 `backdrop-filter: blur(12px)` 毛玻璃效果，hover 时上浮 + 阴影扩散。

## 页面结构

### 首页 `/`

```
┌─────────────────────────────────────────┐
│  Logo    搜索框    [购物车] [用户]       │  ← 毛玻璃导航栏
├─────────────────────────────────────────┤
│                                         │
│    「发现你的数字好物」                    │  ← Hero 区域
│    简单搜索，快速找到你需要的              │     薄荷绿→天蓝渐变背景
│    [搜索框]                              │
│                                         │
├─────────────────────────────────────────┤
│  [全部] [游戏] [会员] [软件] [其他]       │  ← 分类标签栏
├─────────────────────────────────────────┤
│  ┌─────┐  ┌─────┐  ┌─────┐             │
│  │商品1│  │商品2│  │商品3│              │  ← 商品卡片网格
│  │ ¥XX │  │ ¥XX │  │ ¥XX │             │     3列响应式
│  └─────┘  └─────┘  └─────┘             │
│  ┌─────┐  ┌─────┐  ┌─────┐             │
│  │商品4│  │商品5│  │商品6│              │
│  └─────┘  └─────┘  └─────┘             │
├─────────────────────────────────────────┤
│  © 2026 MiKa · 联系方式 · 帮助中心       │  ← 页脚
└─────────────────────────────────────────┘
```

### 商品详情 `/product/:id`

```
┌─────────────────────────────────────────┐
│  导航栏                                  │
├─────────────────────────────────────────┤
│  ┌──────────────┐  商品名称              │
│  │              │  分类标签              │
│  │   商品图片    │  ¥价格                │
│  │              │  库存: XX             │
│  └──────────────┘  数量: [- 1 +]        │
│                    [立即购买]            │
│  商品描述...                             │
│  ─────────────────────                  │
│  用户评价 (0)                            │
└─────────────────────────────────────────┘
```

### 购物车 `/cart`

```
┌─────────────────────────────────────────┐
│  导航栏                                  │
├─────────────────────────────────────────┤
│  购物车 (2)                              │
│  ┌─────────────────────────────────┐    │
│  │ 商品名  x2    ¥XX    [删除]     │    │
│  ├─────────────────────────────────┤    │
│  │ 商品名  x1    ¥XX    [删除]     │    │
│  └─────────────────────────────────┘    │
│  优惠券: [输入码] [验证]                 │
│  ─────────────────────                  │
│  合计: ¥XX                              │
│  [去结算]                               │
└─────────────────────────────────────────┘
```

### 订单 `/order`

```
┌─────────────────────────────────────────┐
│  导航栏                                  │
├─────────────────────────────────────────┤
│  我的订单                                │
│  ┌─────────────────────────────────┐    │
│  │ 订单号: XXXX  状态: 已完成       │    │
│  │ 商品: XXX     ¥XX               │    │
│  │ 卡密: [点击查看]                 │    │
│  └─────────────────────────────────┘    │
│  ┌─────────────────────────────────┐    │
│  │ 订单号: XXXX  状态: 待支付       │    │
│  │ [取消订单] [继续支付]            │    │
│  └─────────────────────────────────┘    │
└─────────────────────────────────────────┘
```

## 后端 API

### 认证

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/auth/register` | 用户注册 |
| POST | `/api/auth/login` | 用户登录 |
| GET | `/api/auth/profile` | 获取当前用户信息 |

### 商品

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/products` | 商品列表（支持 `?category=` 筛选） |
| GET | `/api/products/:id` | 商品详情 |

### 订单

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/orders` | 创建订单（需登录） |
| GET | `/api/orders/:id` | 订单详情 + 卡密（需登录） |
| GET | `/api/user/orders` | 我的订单列表（需登录） |

### 优惠券

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/coupons/verify` | 验证优惠券有效性 |

## 数据模型

### User

```go
type User struct {
    ID           uint      `gorm:"primaryKey"`
    Username     string    `gorm:"uniqueIndex;size:50"`
    Email        string    `gorm:"uniqueIndex;size:100"`
    PasswordHash string    `gorm:"size:255"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
```

### Product

```go
type Product struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `gorm:"size:100"`
    Description string    `gorm:"type:text"`
    Price       float64   `gorm:"type:decimal(10,2)"`
    Category    string    `gorm:"size:50;index"`
    ImageURL    string    `gorm:"size:255"`
    Stock       int       // 库存数量
    Status      string    `gorm:"size:20"` // active/inactive
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### Card

```go
type Card struct {
    ID        uint   `gorm:"primaryKey"`
    ProductID uint   `gorm:"index"`
    Content   string `gorm:"type:text"` // 卡密内容
    OrderID   *uint  `gorm:"index"`     // nil = 未售出
    Status    string `gorm:"size:20"`   // unused/used
    CreatedAt time.Time
}
```

### Order

```go
type Order struct {
    ID           uint      `gorm:"primaryKey"`
    OrderNo      string    `gorm:"uniqueIndex;size:32"`
    UserID       uint      `gorm:"index"`
    ProductID    uint      `gorm:"index"`
    ProductName  string    `gorm:"size:100"`
    Quantity     int
    UnitPrice    float64   `gorm:"type:decimal(10,2)"`
    TotalPrice   float64   `gorm:"type:decimal(10,2)"`
    CouponCode   string    `gorm:"size:50"`
    Discount     float64   `gorm:"type:decimal(10,2)"`
    Status       string    `gorm:"size:20"` // pending/paid/cancelled
    Cards        []Card    `gorm:"foreignKey:OrderID"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
```

### Coupon

```go
type Coupon struct {
    ID            uint      `gorm:"primaryKey"`
    Code          string    `gorm:"uniqueIndex;size:50"`
    DiscountType  string    `gorm:"size:20"` // fixed/percent
    DiscountValue float64   `gorm:"type:decimal(10,2)"`
    MinAmount     float64   `gorm:"type:decimal(10,2)"`
    ExpiresAt     time.Time
    UsageLimit    int
    UsedCount     int
    Status        string    `gorm:"size:20"` // active/inactive
    CreatedAt     time.Time
}
```

## 技术栈

- **前端**: Vue 3 + TypeScript + Vite + Pinia + Vue Router
- **后端**: Go + Gin + GORM + SQLite
- **样式**: 纯 CSS（无 UI 框架），手写清新渐变主题

## 实现范围

### 包含

- 用户注册/登录（JWT 认证）
- 商品展示（列表 + 详情）
- 购物车
- 下单 + 自动发卡
- 订单查询 + 卡密查看
- 优惠券验证
- 后台管理：商品 CRUD、卡密导入、订单管理

### 不包含

- 在线支付集成（预留接口，手动确认支付）
- 评论系统
- 推荐算法
- 多语言
