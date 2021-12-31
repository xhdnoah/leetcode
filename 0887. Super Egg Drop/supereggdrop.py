class Solution:
    def superEggDrop(self, K: int, N: int) -> int:
        memo = dict()

        def dp(k, n):
            if k == 1:
                return n
            if n == 0:
                return 0

            if (k, n) in memo:
                return memo[(k, n)]
            res = float('INF')
            for i in range(1, N+1):
                res = min(
                    res,
                    max(
                        dp(k-1, i-1),
                        dp(k, n-i)
                    ) + 1
                )
            memo[(k, n)] = res
            return res
        return dp(K, N)

    def superEggDropWithBinarySearch(self, K, N) -> int:
        memo = dict()

        def dp(k, n):
            if k == 1:
                return n
            if n == 0:
                return 0

            if (k, n) in memo:
                return memo[(k, n)]

            left = 1
            right = n
            res = float('INF')
            while(left <= right):
                mid = (left + right) // 2
                broken = dp(k-1, mid-1)
                not_broken = dp(k, n-mid)
                if not_broken > broken:
                    left = mid + 1
                    res = min(res, not_broken + 1)
                else:
                    right = mid - 1
                    res = min(res, broken + 1)
            memo[(k, n)] = res
            return res
        return dp(K, N)
