#include <vector>
using namespace std;

class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
        int length = nums.size();
        vector<int> dp(length, 1);
        for (int i = 0; i < length; i++) {
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j])
                    dp[i] = max(dp[i], dp[j] + 1);
            }
        }
        int res = 0;
        for (int i = 0; i < length; i++) {
            res = max(dp[i], res);
        }
        return res;
    }
};
