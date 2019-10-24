# HTMLSimilarity

根据网页结构判断页面相似性(Determine page similarity based on HTML page structure)

## 说明

### 输入参数：
* HTML文档1
* HTML文档2
* 降维后的维数，默认是5000

### 返回值：
* 是否相似
* 相似值（value<0.2时相似，value>0.2时不相似）


## 判断方法

根据网页的DOM树确定网页的模板特征向量，对模板特征向量计算网页结构相似性。

## 引用

> 详细参考：[李景阳, 张波. 网页结构相似性确定方法及装置:.](http://cprs.patentstar.com.cn/Search/Detail?ANE=9HCC7BGA7AHACGEA7GAA8BHA5ADA9FGF8CBA9EDA9BDC9FCG) 原理参考上述专利文章，对其判断相似性部分进行简单实现。

> 本项目是对 [SPuerBRead/HTMLSimilarity](https://github.com/SPuerBRead/HTMLSimilarity) 用 Golang 实现