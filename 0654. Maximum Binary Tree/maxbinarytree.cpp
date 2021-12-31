/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* check(vector<int>& nums, int left, int right) {
        if (left > right) return nullptr;

        int maxIndex = left;
        for (int i = left + 1; i <= right; i++) {
            if (nums[i] > nums[maxIndex]) {
                maxIndex = i;
            }
        }

        TreeNode* node = new TreeNode(nums[maxIndex]);
        node->left = check(nums, left, maxIndex - 1);
        node->right = check(nums, maxIndex+1, right);
        return node;
    }
    TreeNode* constructMaximumBinaryTree(vector<int>& nums) {
        TreeNode* root;
        int right = nums.size() - 1;
        root = check(nums, 0, right);
        return root;
    }
};
