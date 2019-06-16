## 排序算法

### 插入排序
- 描述

    插入排序通过循环方式，将每一个元素放置到前面的有序数组中，放置过程中需要移动目标和当前位置的所有元素。

### 快速排序
- 描述

    快速排序采用分治算法：分解、解决、合并。
    
    **分解：**  根据某个目标值对数组分割成两个子数组A[p..q-1],A[q+1..r]，使得左边数组的每个元素都小于A[q],右边数组的每个元素都大于A[q].
    **解决：**  通过递归调用快速排序，分别对两个子数组进行排序
    **合并：**  最终组合各排序结果。
- 伪代码

<pre><code>
QuickSort(A, p, r)
    if(p < r) {
        q = Partition(A, p, r)
        QuickSort(A, p, q - 1)
        QuickSort(A, q + 1, r)
    }

    
    
Partition(A, p, r)
    pivot = A[r]
    i, j = p - 1, p
    for j = p => r {
        if(A[j] <= pivot) {
            i++
            swap(A[i], A[j])
        }
    }
    swap(A[i + 1], A[r])
    return i + 1