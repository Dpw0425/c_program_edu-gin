# CProgramEdu 数据模型

### 用户表 (`users`)

| 属性         | 数据类型            | 说明    | 备注                 |
|------------|-----------------|-------|--------------------|
| id         | bigint unsigned | 自增 ID | primary key        |
| created_at | datetime        | 创建时间  |                    |
| updated_at | datetime        | 修改时间  |                    |
| deleted_at | datetime        | 删除时间  |                    |
| user_id    | bigint          | 用户 ID | 雪花算法生成 int64 唯一序列号 |
| nick_name  | varchar(20)     | 昵称    |                    |
| password   | varchar(255)    | 密码    |                    |
| email      | varchar(30)     | 邮箱    |                    |
| avatar     | longtext        | 头像    |                    |
| status     | tinyint         | 状态    | 1: 正常; 2: 停用       |

### 题目表 (`questions`)

| 属性           | 数据类型            | 说明     | 备注          |
|--------------|-----------------|--------|-------------|
| id           | bigint unsigned | 自增 ID  | primary key |
| created_at   | time            | 创建时间   |             |
| updated_at   | time            | 修改时间   |             |
| deleted_at   | time            | 删除时间   |             |
| title        | varchar(255)    | 标题     |             |
| content      | longtext        | 题目内容   |             |
| tag          | longtext        | 标签     |             |
| degree       | tinyint         | 难度     |             |
| owner_id     | bigint          | 发布人 ID |             |
| passing_rate | double          | 正确率    |             |

### 题目测试数据表 (`test_data`)

| 属性          | 数据类型            | 说明    | 备注          |
|-------------|-----------------|-------|-------------|
| id          | bigint unsigned | 自增 ID | primary key |
| created_at  | time            | 创建时间  |             |
| updated_at  | time            | 修改时间  |             |
| deleted_at  | time            | 删除时间  |             |
| question_id | bigint unsigned | 题目 ID |             |
| input       | longtext        | 测试用例  |             |
| output      | longtext        | 期望输出  |             |

### 用户题目关系表 (`user_ques`)

| 属性             | 数据类型            | 说明    | 备注                                           |
|----------------|-----------------|-------|----------------------------------------------|
| id             | bigint unsigned | 自增 ID | primary key                                  |
| created_at     | time            | 创建时间  |                                              |
| updated_at     | time            | 修改时间  |                                              |
| deleted_at     | time            | 删除时间  |                                              |
| user_id        | bigint          | 用户 ID |                                              |
| question_id    | bigint unsigned | 题目 ID |                                              |
| competition_id | bigint unsigned | 比赛 ID | omit empty                                   |
| status         | varchar(20)     | 完成情况  | default: incomplete, oneof"failed, accepted" |

### 比赛表 (`competitions`)

| 属性         | 数据类型            | 说明     | 备注                          |
|------------|-----------------|--------|-----------------------------|
| id         | bigint unsigned | 自增 ID  | primary key                 |
| created_at | time            | 创建时间   |                             |
| updated_at | time            | 修改时间   |                             |
| deleted_at | time            | 删除时间   |                             |
| name       | longtext        | 比赛名称   |                             |
| contestant | longtext        | 参赛选手   |                             |
| owner_id   | bigint          | 创办者 ID |                             |
| start_time | time            | 开始时间   |                             |
| deadline   | time            | 截止时间   |                             |
| status     | tinyint         | 比赛状态   | 0: 报名中 1: 未开始 2: 进行中 3: 已截止 |
| category   | tinyint         | 比赛类别   | 0: 个人 1: 团体                 |
| permission | tinyint         | 比赛权限   | 0: 公开(所有人可报名) 1: 私有(报名需审批)  |

### 比赛报名表 (`entries`)

| 属性             | 数据类型            | 说明    | 备注                  |
|----------------|-----------------|-------|---------------------|
| id             | bigint unsigned | 自增 ID | primary key         |
| created_at     | time            | 创建时间  |                     |
| updated_at     | time            | 修改时间  |                     |
| deleted_at     | time            | 删除时间  |                     |
| user_id        | bigint          | 用户 ID |                     |
| competition_id | bigint unsigned | 比赛 ID |                     |
| status         | tinyint         | 审核状态  | 0: 审核中 1: 通过 2: 未通过 |
