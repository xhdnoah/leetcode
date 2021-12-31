class Solution:
  def min_distance(self, word1: str, word2: str) -> int:
    memo = dict()

    def dp(i, j):
      if i == -1: return j+1
      if j == -1: return i+1
      if (i, j) in memo:
        return memo[(i, j)]
      if word1[i] == word2[j]:
        memo[(i, j)] = dp(i-1, j-1)
      else:
        memo[(i, j)] = min(
        dp(i, j-1) + 1,  # insert
        dp(i-1, j) + 1,  # delete
        dp(i-1, j-1) + 1  # replace
      )
      return memo[(i, j)]
    return dp(len(word1) - 1, len(word2) - 1)
