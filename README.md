# ALGO
![](https://img.shields.io/badge/language-golang-blue.svg)
[![Forks](https://img.shields.io/github/forks/rosemary666/algo)](https://img.shields.io/github/forks/rosemary666/algo)
[![Stars](https://img.shields.io/github/stars/rosemary666/algo)](https://img.shields.io/github/stars/rosemary666/algo)
[![Gitbook Action Build](https://github.com/rosemary666/algo/actions/workflows/gitbook_action.yaml/badge.svg)](https://github.com/rosemary666/algo/actions/workflows/gitbook_action.yaml)
[![License: Apache2.0](https://img.shields.io/github/license/rosemary666/algo)](https://github.com/rosemary666/algo/blob/main/LICENSE)

[在线DOCS](https://rosemary666.github.io/algo/)

主要使用golang语言讲解和实现算法，包括数据结构、leetcode算法题等。   
Leetcode部分主要参考了[代码随想录](https://programmercarl.com/), 然后自己做了相关的笔记和实现，便于分享和后期温习。

## 相关题目题目

### 数组
| 题目 | 原题 | 难度 |问题和讲解 | 源码 | 关键词 | 已完成 |
| :-----| ----: | ----: |:----: | :----: | :----: |:----: |
| 704. 二分查找 | [leecode原题链接](https://leetcode.cn/binary-search/) | 简单 |[问题和讲解](./array/704-二分查找.md) | [源码](./array/code/704-binary-search/main.go) | 二分法 |&#9745; |
| 27. 移除元素 | [leecode原题链接](https://leetcode.cn/remove-element/) | 简单 | [问题和讲解](./array/27-移除元素.md) | [源码](./array/code/27-remove-element/main.go) | 双指针法 |&#9745; |
| 977. 有序数组的平方 | [leecode原题链接](https://leetcode.cn/squares-of-a-sorted-array/) | 简单 |[问题和讲解](./array/977-有序数组的平方.md) | [源码](./array/code/977-squares-of-a-sorted-array/main.go) | 双指针法 |&#9745; |
| 209. 长度最小的子数组 | [leecode原题链接](https://leetcode.cn/minimum-size-subarray-sum/) | 中等 | [问题和讲解](./array/209-长度最小的子数组.md) | [源码](./array/code/209-minimum-size-subarray-sum/main.go) | 滑动窗口 |&#9745; |
| 59. 螺旋矩阵 II | [leecode原题链接](https://leetcode.cn/spiral-matrix-ii/) | 中等 | [问题和讲解](./array/59-螺旋矩阵.md) | [源码](./array/code/59-spiral-matrix-ii/main.go) | 考验代码掌控能力和边界掌控 |&#9745; |

### 链表
| 题目 | 原题 | 难度 |问题和讲解 | 源码 | 关键词 |已完成 |
| :-----| ----: | ----: |:----: | :----: | :----: |:----: |
| 203. 移除链表元素 | [leecode原题链接](https://leetcode.cn/remove-linked-list-elements/) | 简单 |[问题和讲解](./linked-list/203-移除链表元素.md) | [源码](./linked-list/code/203-remove-linked-list-elements/main.go) | 虚拟头节点 | &#9745; |
| 707. 设计链表 | [leecode原题链接](https://leetcode.cn/design-linked-list/) | 中等 |[问题和讲解](./linked-list/707-设计链表.md) | [源码](./linked-list/code/707-design-linked-list/main.go) | 链表基本操作 | &#9745; |
| 206. 反转链表 | [leecode原题链接](https://leetcode.cn/reverse-linked-list/) | 简单 |[问题和讲解](./linked-list/206-反转链表) | [源码](./linked-list/code/206-reverse-linked-list/main.go) | | &#9745; |
| 24. 两两交换链表中的节点 | [leecode原题链接](https://leetcode.cn/swap-nodes-in-pairs/) | 中等 |[问题和讲解](./linked-list/24-两两交换链表中的节点.md) | [源码](./linked-list/code/24-swap-nodes-in-pairs/main.go) | | &#9745; |
| 19. 删除链表的倒数第 N 个结点 | [leecode原题链接](https://leetcode.cn/remove-nth-node-from-end-of-list/) | 中等 |[问题和讲解](./linked-list/19-删除链表的倒数第N个节点.md) | [源码](./linked-list/code/19-remove-nth-node-from-end-of-list/main.go) | 双指针法 | &#9745; |
| 面试题 02.07. 链表相交 | [leecode原题链接](https://leetcode.cn/intersection-of-two-linked-lists-lcci/) | 简单 |[问题和讲解](./linked-list/面试题-0207-链表相交.md) | [源码](./linked-list/code/0207-intersection-of-two-linked-lists-lcci/main.go) | | &#9745; |
| 142. 环形链表 II | [leecode原题链接](https://leetcode.cn/linked-list-cycle-ii/) | 中等 |[问题和讲解](./linked-list/142-环形链表II.md) | [源码](./linked-list/code/142-linked-list-cycle-ii/main.go) | | &#9745; |

### 哈希表
| 题目 | 原题 | 难度 |问题和讲解 | 源码 | 关键词 |已完成 |
| :-----| ----: | ----: |:----: | :----: | :----: |:----: |
| 242. 有效的字母异位词 | [leecode原题链接](https://leetcode.cn/valid-anagram/) | 简单 |[问题和讲解](./hash/242-有效的字母异位词.md) | [源码](./hash/code/242-valid-anagram/main.go) | |&#9745; |
| 349. 两个数组的交集 | [leecode原题链接](https://leetcode.cn/intersection-of-two-arrays/) | 简单 |[问题和讲解](./hash/349-两个数组的交集.md) | [源码](./hash/code/349-intersection-of-two-arrays/main.go) | |&#9745; |
| 202. 快乐数 | [leecode原题链接](https://leetcode.cn/happy-number/) | 简单 |[问题和讲解](./hash/202-快乐数.md) | [源码](./hash/code/202-happy-number/main.go) | |&#9745; |
| 1. 两数之和 | [leecode原题链接](https://leetcode.cn/two-sum/) | 简单 |[问题和讲解](./hash/1-两数之和.md) | [源码](./hash/code/1-two-sum/main.go) | |&#9745; |
| 454. 四数相加 II | [leecode原题链接](https://leetcode.cn/4sum-ii/) | 中等 |[问题和讲解](./hash/454-四数相加II.md) | [源码](./hash/code/454-4sum-ii/main.go) | |&#9745; |
| 383. 赎金信 | [leecode原题链接](https://leetcode.cn/ransom-note/) | 简单 |[问题和讲解](./hash/383-赎金信.md) | [源码](./hash/code/383-ransom-note/main.go) | |&#9745; |
| 15. 三数之和 | [leecode原题链接](https://leetcode.cn/3sum/) | 中等 |[问题和讲解](./hash/15-三数之和.md) | [源码](./hash/code/15-3sum/main.go) | |&#9745; |
| 18. 四数之和 | [leecode原题链接](https://leetcode.cn/4sum/) | 中等 |[问题和讲解](./hash/18-四数之和.md) | [源码](./hash/code/18-4sum/main.go) | |&#9745; |

### 字符串
| 题目 | 原题 | 难度 |问题和讲解 | 源码 | 关键词 |已完成 |
| :-----| ----: | ----: |:----: | :----: | :----: |:----: |
| 344. 反转字符串 | [leecode原题链接](https://leetcode.cn/reverse-string/) | 简单 |[问题和讲解](./string/344-反转字符串.md) | [源码](./string/code/344-reverse-string/main.go) | |&#9745; |
| 541. 反转字符串 II | [leecode原题链接](https://leetcode.cn/reverse-string-ii/) | 简单 |[问题和讲解](./string/541-反转字符串II.md) | [源码](./string/code/541-reverse-string-ii/main.go) | |&#9745; |
| 剑指 Offer 05. 替换空格 | [leecode原题链接](https://leetcode.cn/ti-huan-kong-ge-lcof/) | 简单 |[问题和讲解](./string/541-反转字符串II.md) | [源码](./string/code/05-ti-huan-kong-ge-lcof/main.go) | |&#9745; |
| 151. 颠倒字符串中的单词 | [leecode原题链接](https://leetcode.cn/reverse-words-in-a-string/) | 中等 |[问题和讲解](./string/151-颠倒字符串中的单词.md) | [源码](./string/code/151-reverse-words-in-a-string/main.go) | |&#9745; |
| 58 - II. 左旋转字符串 | [leecode原题链接](https://leetcode.cn/zuo-xuan-zhuan-zi-fu-chuan-lcof/) | 简单 |[问题和讲解](./string/58-II左旋转字符串.md) | [源码](./string/code/58-zuo-xuan-zhuan-zi-fu-chuan-lcof/main.go) | 花式反转 |&#9745; |
| 28. 实现 strStr | [leecode原题链接](https://leetcode.cn/implement-strstr/) | 简单 |[问题和讲解](./string/28-实现strStr.md) | [源码](./string/code/28-implement-strstr/main.go) | KMP算法 |&#9745; |
| 459. 重复的子字符串 | [leecode原题链接](https://leetcode.cn/repeated-substring-pattern/) | 简单 |[问题和讲解](./string/459-重复的子字符串.md) | [源码](./string/code/459-repeated-substring-pattern/main.go) | KMP算法 |&#9745; |

### 栈和队列
| 题目 | 原题 | 难度 |问题和讲解 | 源码 | 关键词 | 已完成 |
| :-----| ----: | ----: |:----: | :----: | :----: |:----: |
| 232. 用栈实现队列 | [leecode原题链接](https://leetcode.cn/implement-queue-using-stacks/) | 简单 |[问题和讲解](./stacks-queues/232-用栈实现队列.md) | [源码](./stacks-queues/code/232-implement-queue-using-stacks/main.go) | 输入、输出栈 |&#9745; |
| 225. 用队列实现栈 | [leecode原题链接](https://leetcode.cn/implement-stack-using-queues/) | 简单 |[问题和讲解](./stacks-queues/225-用队列实现栈.md) | [源码](./stacks-queues/code/225-implement-stack-using-queues/main.go) |  队列的基本使用 |&#9745; |
| 20. 有效的括号 | [leecode原题链接](https://leetcode.cn/valid-parentheses/) | 简单 |[问题和讲解](./stacks-queues/20-有效的括号.md) | [源码](./stacks-queues/code/20-valid-parentheses/main.go) |使用栈解决的经典问题 | &#9745; |
| 1047. 删除字符串中的所有相邻重复项 | [leecode原题链接](https://leetcode.cn/remove-all-adjacent-duplicates-in-string/) | 简单 |[问题和讲解](./stacks-queues/1047-删除字符串中的所有相邻重复项.md) | [源码](./stacks-queues/code/1047-remove-all-adjacent-duplicates-in-string/main.go) | 匹配问题也使用栈解决 |&#9745; |
| 150. 逆波兰表达式求值| [leecode原题链接](https://leetcode.cn/evaluate-reverse-polish-notation/) | 中等 |[问题和讲解](./stacks-queues/150-逆波兰表达式求值.md) | [源码](./stacks-queues/code/150-evaluate-reverse-polish-notation/main.go) | 计算机运算思考方式，使用栈解决的经典问题 |&#9745; |
| 239. 滑动窗口最大值| [leecode原题链接](https://leetcode.cn/sliding-window-maximum/) | 困难 |[问题和讲解](./stacks-queues/239-滑动窗口最大值.md) | [源码](./stacks-queues/code/239-sliding-window-maximum/main.go) | 单调队列的经典题目 |&#9745; |
| 347. 前 K 个高频元素| [leecode原题链接](https://leetcode.cn/top-k-frequent-elements/) | 中等 |[问题和讲解](./stacks-queues/347-前K个高频元素.md) | [源码](./stacks-queues/code/347-top-k-frequent-elements/main.go) | 优先级队列 |&#9745; |

### 二叉树
| 题目 | 原题 | 难度 |问题和讲解 | 源码 | 关键词 | 已完成 |
| :-----| ----: | ----: |:----: | :----: | :----: |:----: |
| 144. 二叉树的前序遍历 | [leecode原题链接](https://leetcode.cn/problems/binary-tree-preorder-traversal/) | 简单 |[问题和讲解](./tree/144-二叉树的前序遍历.md) | [源码](./tree/code/144-binary-tree-preorder-traversal/main.go) | 前序遍历(递归法+迭代法(借助栈)) |&#9745; |
| 145. 二叉树的后序遍历 | [leecode原题链接](https://leetcode.cn/problems/binary-tree-postorder-traversal/) | 简单 |[问题和讲解](./tree/145-二叉树的后序遍历.md) | [源码](./tree/code/145-binary-tree-postorder-traversal/main.go) | 后序遍历(递归法+迭代法(借助栈)) |&#9745; |
| 94. 二叉树的中序遍历 | [leecode原题链接](https://leetcode.cn/problems/binary-tree-inorder-traversal/) | 简单 |[问题和讲解](./tree/94-二叉树的中序遍历.md) | [源码](./tree/code/94-binary-tree-inorder-traversal/main.go) | 后序遍历(递归法+迭代法(借助栈)) |&#9745; |
| 102. 二叉树的层序遍历 | [leecode原题链接](https://leetcode.cn/problems/binary-tree-level-order-traversal/) | 中等 |[问题和讲解](./tree/102-二叉树的层序遍历.md) | [源码](./tree/code/102-binary-tree-level-order-traversal/main.go) | 层次遍历(借助队列) |&#9745; |
| 226. 翻转二叉树 | [leecode原题链接](https://leetcode.cn/problems/invert-binary-tree/) | 简单 |[问题和讲解](./tree/226-翻转二叉树.md) | [源码](./tree/code/226-invert-binary-tree/main.go) | 可以是层次遍历的的变种实现 |&#9745; |
| 101. 对称二叉树 | [leecode原题链接](https://leetcode.cn/problems/symmetric-tree/) | 简单 |[问题和讲解](./tree/101-对称二叉树.md) | [源码](./tree/code/101-symmetric-tree/main.go) | 递归和迭代  |&#9745; |
| 104. 二叉树的最大深度 | [leecode原题链接](https://leetcode.cn/problems/maximum-depth-of-binary-tree/) | 简单 |[问题和讲解](./tree/104-二叉树的最大深度.md) | [源码](./tree/code/104-maximum-depth-of-binary-tree/main.go) | 层序遍历的用法  |&#9745; |
| 111. 二叉树的最小深度 | [leecode原题链接](https://leetcode.cn/problems/minimum-depth-of-binary-tree/) | 简单 |[问题和讲解](./tree/111-二叉树的最小深度.md) | [源码](./tree/code/111-minimum-depth-of-binary-tree/main.go) | 层序遍历的用法  |&#9745; |
| 222. 完全二叉树的节点个数 | [leecode原题链接](https://leetcode.cn/problems/count-complete-tree-nodes/) | 中等 |[问题和讲解](./tree/222-完全二叉树的节点个数.md) | [源码](./tree/code/222-count-complete-tree-nodes/main.go) |  层次遍历 |&#9745; |
| 110. 平衡二叉树 | [leecode原题链接](https://leetcode.cn/problems/balanced-binary-tree/) | 简单 |[问题和讲解](./tree/110-平衡二叉树.md) | [源码](./tree/code/110-balanced-binary-tree/main.go) |  递归 |&#9745; |
| 257. 二叉树的所有路径 | [leecode原题链接](https://leetcode.cn/problems/binary-tree-paths/) | 简单 |[问题和讲解](./tree/257-二叉树的所有路径.md) | [源码](./tree/code/257-binary-tree-paths/main.go) |  递归和回溯 |&#9745; |
| 404. 左叶子之和 | [leecode原题链接](https://leetcode.cn/problems/sum-of-left-leaves/) | 简单 |[问题和讲解](./tree/404-左叶子之和.md) | [源码](./tree/code/404-sum-of-left-leaves/main.go) |  前序遍历变种 |&#9745; |
| 513. 找树左下角的值 | [leecode原题链接](https://leetcode.cn/problems/find-bottom-left-tree-value/) | 中等 |[问题和讲解](./tree/513-找树左下角的值.md) | [源码](./tree/code/513-find-bottom-left-tree-value/main.go) | 层次遍历的应用  |&#9745; |
| 112. 路径总和 | [leecode原题链接](https://leetcode.cn/problems/path-sum/) | 简单 |[问题和讲解](./tree/112-路径总和.md) | [源码](./tree/code/112-path-sum/main.go) | 跟257这题思路类似  |&#9745; |
| 106. 从中序与后序遍历序列构造二叉树 | [leecode原题链接](https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/) | 中等 |[问题和讲解](./tree/106-从中序与后序遍历序列构造二叉树.md) | [源码](./tree/code/106-construct-binary-tree-from-inorder-and-postorder-traversal/main.go) | 典型的从遍历反向构造二叉树   |&#9745; |
| 654. 最大二叉树 | [leecode原题链接](https://leetcode.cn/problems/maximum-binary-tree/) | 中等 |[问题和讲解](./tree/654-最大二叉树.md) | [源码](./tree/code/654-maximum-binary-tree/main.go) | 递归  |&#9745; |
| 617. 合并二叉树 | [leecode原题链接](https://leetcode.cn/problems/merge-two-binary-trees/) | 简单 |[问题和讲解](./tree/617-合并二叉树.md) | [源码](./tree/code/617-merge-two-binary-trees/main.go) | 两棵子树同时前序遍历  |&#9745; |
| 700. 二叉搜索树中的搜索 | [leecode原题链接](https://leetcode.cn/problems/search-in-a-binary-search-tree/) | 简单 |[问题和讲解](./tree/700-二叉搜索树中的搜索.md) | [源码](./tree/code/700-search-in-a-binary-search-tree/main.go) | 二叉搜索树特性  |&#9745; |
| 98. 验证二叉搜索树 | [leecode原题链接](https://leetcode.cn/problems/validate-binary-search-tree/) | 中等 |[问题和讲解](./tree/98-验证二叉搜索树.md) | [源码](./tree/code/98-validate-binary-search-tree/main.go) | 二叉搜索树特性  |&#9745; |
| 530. 二叉搜索树的最小绝对差 | [leecode原题链接](https://leetcode.cn/problems/minimum-absolute-difference-in-bst/) | 简单 |[问题和讲解](./tree/530-二叉搜索树的最小绝对差.md) | [源码](./tree/code/530-minimum-absolute-difference-in-bst/main.go) | 二叉搜索树特性和中序遍历的结合(形成有序数组)  |&#9745; |
| 501. 二叉搜索树中的众数 | [leecode原题链接](https://leetcode.cn/problems/find-mode-in-binary-search-tree/) | 简单 |[问题和讲解](./tree/501-二叉搜索树中的众数.md) | [源码](./tree/code/501-find-mode-in-binary-search-tree/main.go) | 二叉搜索树特性和中序遍历的结合(进阶: 一次遍历中解决)  |&#9745; |
| 236. 二叉树的最近公共祖先 | [leecode原题链接](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/) | 简单 |[问题和讲解](./tree/236-二叉树的最近公共祖先.md) | [源码](./tree/code/236-lowest-common-ancestor-of-a-binary-tree/main.go) | 递归   |&#9745; |
| 701. 二叉搜索树中的插入操作 | [leecode原题链接](https://leetcode.cn/problems/insert-into-a-binary-search-tree/) | 中等 |[问题和讲解](./tree/701-二叉搜索树中的插入操作.md) | [源码](./tree/code/701-insert-into-a-binary-search-tree/main.go) | 二叉树特性 |&#9745; |
| 450. 删除二叉搜索树中的节点 | [leecode原题链接](https://leetcode.cn/problems/delete-node-in-a-bst/) | 中等 |[问题和讲解](./tree/450-删除二叉搜索树中的节点.md) | [源码](./tree/code/450-delete-node-in-a-bst/main.go) | 二叉搜索树特性利用  |&#9745; |
| 669. 修剪二叉搜索树 | [leecode原题链接](https://leetcode.cn/problems/trim-a-binary-search-tree/) | 中等 |[问题和讲解](./tree/669-修剪二叉搜索树.md) | [源码](./tree/code/669-trim-a-binary-search-tree/main.go) | 二叉搜索树特性利用  |&#9745; |
| 108. 将有序数组转换为二叉搜索树 | [leecode原题链接](https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/) | 简单 |[问题和讲解](./tree/108-将有序数组转换为二叉搜索树.md) | [源码](./tree/code/108-convert-sorted-array-to-binary-search-tree/main.go) | 结合二分法的思想使用  |&#9745; |
| 538. 把二叉搜索树转换为累加树 | [leecode原题链接](https://leetcode.cn/problems/convert-bst-to-greater-tree/) | 中等 |[问题和讲解](./tree/538-把二叉搜索树转换为累加树.md) | [源码](./tree/code/538-convert-bst-to-greater-tree/main.go) | 二叉树的反中序遍历  |&#9745; |

### 回溯算法
| 题目 | 原题 | 难度 |问题和讲解 | 源码 | 关键词 | 已完成 |
| :-----| ----: | ----: |:----: | :----: | :----: |:----: |
| 77. 组合 | [leecode原题链接](https://leetcode.cn/problems/combinations/) | 中等 |[问题和讲解](./backtracking/77-组合.md) | [源码](./backtracking/code/77-combinations/main.go) | 组合类回溯算法  |&#9745; |
| 216. 组合总和 III | [leecode原题链接](https://leetcode.cn/problems/combination-sum-iii/) | 中等 |[问题和讲解](./backtracking/216-组合总和III.md) | [源码](./backtracking/code/216-combination-sum-iii/main.go) | 组合类回溯算法  |&#9745; |
| 17. 电话号码的字母组合 | [leecode原题链接](https://leetcode.cn/problems/letter-combinations-of-a-phone-number/) | 中等 |[问题和讲解](./backtracking/17-电话号码的字母组合.md) | [源码](./backtracking/code/17-letter-combinations-of-a-phone-number/main.go) | 组合类回溯算法  |&#9745; |
| 39. 组合总和 | [leecode原题链接](https://leetcode.cn/problems/combination-sum/) | 中等 |[问题和讲解](./backtracking/39-组合总和.md) | [源码](./backtracking/code/39-combination-sum/main.go) | 组合类回溯算法  |&#9745; |
| 40. 组合总和 II | [leecode原题链接](https://leetcode.cn/problems/combination-sum-ii/) | 中等 |[问题和讲解](./backtracking/40-组合总和II.md) | [源码](./backtracking/code/40-combination-sum-ii/main.go) | 组合类回溯算法（去重）  |&#9745; |
| 131. 分割回文串 | [leecode原题链接](https://leetcode.cn/problems/palindrome-partitioning/) | 中等 |[问题和讲解](./backtracking/131-分割回文串.md) | [源码](./backtracking/code/131-palindrome-partitioning/main.go) | 分割问题  |&#9745; |
| 93. 复原 IP 地址 | [leecode原题链接](https://leetcode.cn/problems/restore-ip-addresses/) | 中等 |[问题和讲解](./backtracking/93-复原IP地址.md) | [源码](./backtracking/code/93-restore-ip-addresses/main.go) | 分割问题  |&#9745; |
| 78. 子集 | [leecode原题链接](https://leetcode.cn/problems/subsets/) | 中等 |[问题和讲解](./backtracking/78-子集.md) | [源码](./backtracking/code/78-subsets/main.go) | 子集问题  |&#9745; |
| 90. 子集 II | [leecode原题链接](https://leetcode.cn/problems/subsets-ii/) | 中等 |[问题和讲解](./backtracking/90-子集II.md) | [源码](./backtracking/code/90-subsets-ii/main.go) | 子集问题  |&#9745; |
| 491. 递增子序列 | [leecode原题链接](https://leetcode.cn/problems/increasing-subsequences/) | 中等 |[问题和讲解](./backtracking/491-递增子序列.md) | [源码](./backtracking/code/491-increasing-subsequences/main.go) | 像子集问题  |&#9745; |
| 46. 全排列 | [leecode原题链接](https://leetcode.cn/problems/permutations/) | 中等 |[问题和讲解](./backtracking/46-全排列.md) | [源码](./backtracking/code/46-permutations/main.go) | 像子集问题   |&#9745; |
| 47. 全排列 II | [leecode原题链接](https://leetcode.cn/problems/permutations-ii/) | 中等 |[问题和讲解](./backtracking/47-全排列II.md) | [源码](./backtracking/code/47-permutations-ii/main.go) | 像子集问题（去重）   |&#9745; |
| 332. 重新安排行程 | [leecode原题链接](https://leetcode.cn/problems/reconstruct-itinerary/) | 困难 |[问题和讲解](./backtracking/332-重新安排行程.md) | [源码](./backtracking/code/332-reconstruct-itinerary/main.go) |    |&#9745; |
| 51. N 皇后 | [leecode原题链接](https://leetcode.cn/problems/n-queens/) | 困难 |[问题和讲解](./backtracking/51-N皇后.md) | [源码](./backtracking/code/51-n-queens/main.go) |    |&#9745; |
| 37. 解数独| [leecode原题链接](https://leetcode.cn/problems/sudoku-solver/) | 困难 |[问题和讲解](./backtracking/37-解数独.md) | [源码](./backtracking/code/37-sudoku-solver/main.go) |    |&#9745; |