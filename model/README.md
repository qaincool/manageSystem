role 角色表
分为：
    1. 管理员：admin
    2. 学生: student
    3. 教师: teacher

permission 权限表
分为：
    1. 用户增加
    2. 用户删除
    3. 用户修改
    4. 用户查询
    5. 视频增加
    6. 视频删除
    7. 视频修改
    8. 视频查询

角色和权限的关系
管理员：所有权限
学生：只有视频的查询权限
教师：有视频的所有权限


1	user	用户管理		1
10	user_add	用户增加	1	2
11	user_del	用户删除	1	3
12	user_edit	用户修改	1	4
13	user_select	用户查询	1	5
2	video	视频管理		6
21	video_del	视频删除	2	7
22	video_edit	视频修改	2	8
20	video_add	视频增加	2	9
23	video_select	视频查询	2	10