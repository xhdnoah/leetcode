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
    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        return build(preorder, 0, preorder.size() - 1,
                     inorder, 0, inorder.size() - 1);
    }

    TreeNode* build(vector<int>& preorder, int preStart, int preEnd,
                    vector<int>& inorder, int inStart, int inEnd){
        if (preStart > preEnd) {
            return nullptr;
        }

        TreeNode* root = new TreeNode(preorder[preStart]);

        int index;
        for (int i = 0; i < inorder.size(); i++) {
            if (inorder[i] == root->val) {
                index = i;
            }
        }

        int leftSize = index - inStart;

        root->left = build(preorder, preStart + 1, preStart + leftSize,
                          inorder, inStart, index - 1);
        root->right = build(preorder, preStart + leftSize + 1, preEnd,
                           inorder, index + 1, inEnd);

        return root;
    }

};
