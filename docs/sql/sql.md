# CProgramEdu 数据模型

### 用户表 (`users`)

| 属性         | 数据类型            | 说明    | 备注                 |
|------------|-----------------|-------|--------------------|
| id         | bigint unsigned | 自增 ID | primary key        |
| created_at | datetime        | 创建时间  |                    |
| updated_at | datetime        | 修改时间  |                    |
| deleted_at | datetime        | 删除时间  |                    |
| user_id    | bigint          | 用户 ID | 雪花算法生成 int64 唯一序列号 |
| user_name  | varchar(20)     | 姓名    |                    |
| student_id | bigint unsigned | 学号    | unique             |
| password   | varchar(255)    | 密码    |                    |
| email      | varchar(30)     | 邮箱    |                    |
| avatar     | longtext        | 头像    |                    |
| grade      | tinyint         | 所属年级  |                    |
| status     | tinyint         | 状态    | 1: 正常; 2: 停用       |

### 管理员表 (`admin`)

| 属性         | 数据类型            | 说明    | 备注           |
|------------|-----------------|-------|--------------|
| id         | bigint unsigned | 自增 ID | primary key  |
| created_at | datetime        | 创建时间  |              |
| updated_at | datetime        | 修改时间  |              |
| deleted_at | datetime        | 删除时间  |              |
| username   | varchar(20)     | 用户名   |              |
| teacher_id | bigint unsigned | 教师工号  |              |
| password   | varchar(255)    | 密码    |              |
| permission | varchar(255)    | 权限    | 暂未确定具体值      |
| status     | tinyint         | 状态    | 1: 正常; 2: 停用 |

### 题目表 (`questions`)

| 属性           | 数据类型            | 说明     | 备注          |
|--------------|-----------------|--------|-------------|
| id           | bigint unsigned | 自增 ID  | primary key |
| created_at   | time            | 创建时间   |             |
| updated_at   | time            | 修改时间   |             |
| deleted_at   | time            | 删除时间   |             |
| title        | varchar(255)    | 标题     |             |
| content      | longtext        | 题目内容   |             |
| answer       | longtext        | 参考答案   |             |
| tag          | longtext        | 标签     |             |
| degree       | tinyint         | 难度     |             |
| owner_id     | bigint          | 发布人 ID |             |
| passing_rate | double          | 正确率    |             |
| status       | tinyint         | 题目状态   | 0: 公开 1: 私密 |

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

### 题库表 (`question_bank`)

| 属性           | 数据类型            | 说明    | 备注          |
|--------------|-----------------|-------|-------------|
| id           | bigint unsigned | 自增 ID | primary key |
| created_at   | time            | 创建时间  |             |
| updated_at   | time            | 修改时间  |             |
| deleted_at   | time            | 删除时间  |             |
| name         | varchar(255)    | 题库名称  |             |
| content      | longtext        | 内容    | 包含的题目       |
| open_grade   | longtext        | 开放年级  |             |
| participants | bigint unsigned | 参与人数  |             |
| completed    | bigint unsigned | 完成人数  |             |

### 题库题目关系表 (`bank_ques`)

| 属性           | 数据类型            | 说明    | 备注          |
|--------------|-----------------|-------|-------------|
| id           | bigint unsigned | 自增 ID | primary key |
| created_at   | time            | 创建时间  |             |
| updated_at   | time            | 修改时间  |             |
| deleted_at   | time            | 删除时间  |             |
| bank_id      | bigint unsigned | 题库 ID |             |
| question_id  | bigint unsigned | 题目 ID |             |
| commit_times | bigint unsigned | 提交次数  |             |
| accepted     | bigint unsigned | 通过人数  |             |

### 用户题目关系表 (`user_ques`)

| 属性           | 数据类型            | 说明    | 备注                                           |
|--------------|-----------------|-------|----------------------------------------------|
| id           | bigint unsigned | 自增 ID | primary key                                  |
| created_at   | time            | 创建时间  |                                              |
| updated_at   | time            | 修改时间  |                                              |
| deleted_at   | time            | 删除时间  |                                              |
| user_id      | bigint          | 用户 ID |                                              |
| question_id  | bigint unsigned | 题目 ID |                                              |
| commit_times | bigint unsigned | 提交次数  |                                              |
| status       | varchar(20)     | 完成情况  | default: incomplete, oneof"failed, accepted" |

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

### 比赛题目关系表 (`cpt_ques`)

| 属性             | 数据类型            | 说明    | 备注          |
|----------------|-----------------|-------|-------------|
| id             | bigint unsigned | 自增 ID | primary key |
| created_at     | time            | 创建时间  |             |
| updated_at     | time            | 修改时间  |             |
| deleted_at     | time            | 删除时间  |             |
| competition_id | bigint unsigned | 比赛 ID |             |
| question_id    | bigint unsigned | 题目 ID |             |
| commit_times   | bigint unsigned | 提交次数  |             |
| accepted       | bigint unsigned | 通过人数  |             |

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

### 团队表 (`teams`)

| 属性                | 数据类型            | 说明     | 备注          |
|-------------------|-----------------|--------|-------------|
| id                | bigint unsigned | 自增 ID  | primary key |
| created_at        | time            | 创建时间   |             |
| updated_at        | time            | 修改时间   |             |
| deleted_at        | time            | 删除时间   |             |
| name              | longtext        | 团队名称   |             |
| manager           | bigint          | 队长 ID  |             |
| member            | longtext        | 团队成员   |             |
| competition_times | bigint unsigned | 参与比赛次数 |             |
