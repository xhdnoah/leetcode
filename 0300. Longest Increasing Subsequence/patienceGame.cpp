#include <vector>
using namespace std;

class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
        vector<int> top(nums.size());
        int piles = 0;
        for (int i = 0; i < nums.size(); i++) {
            int left = 0;
            int right = piles;
            int poker = nums[i];
            while(left < right) {
                int mid = (left + right) / 2;
                if (top[mid] >= poker) {
                    right = mid;
                } else {
                    left = mid + 1;
                }
            }
            if (left == piles) piles++;
            top[left] = poker;
        }
        return piles;
    }
};
