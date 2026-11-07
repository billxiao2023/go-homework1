# MetaNode Academy Go 作业提交仓库

每个章节的作业对应一个文件夹。仓库根目录的 GitHub Actions 会扫描所有含 `go.mod` 的文件夹并自动运行测试评分。

## 作业列表

| 章节 | 目录 | 题目 |
| --- | --- | --- |
| 作业 01 | [homework01/](homework01/) | 只出现一次的数字、回文数、有效的括号、最长公共前缀、加一、删除有序数组中的重复项、合并区间、两数之和 |

## 评分

推送后 GitHub Actions 自动执行评分，结果见仓库的 **Actions** 页面。也可在本地对应目录运行：

```bash
cd homework01
go test -v ./...
```
