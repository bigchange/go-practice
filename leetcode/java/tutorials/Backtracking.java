package com.bigchange.algorithm.leetcode.tutorials;

/**
 * User: JerryYou
 *
 * Date: 2019-06-19
 *
 * Copyright (c) 2018 devops
 *
 * <<licensetext>>
 */
public class Backtracking {

  /*
  回溯算法框架：
   vector<int> t;
   void dfs(int cur, int n) {
      if (cur == n) {
          // 记录答案
          // ...
          return;
      }
      // 考虑选择当前位置
      t.push_back(cur);
      dfs(cur + 1, n, k);
      t.pop_back();
      // 考虑不选择当前位置
      dfs(cur + 1, n, k);
   }
  */
}
