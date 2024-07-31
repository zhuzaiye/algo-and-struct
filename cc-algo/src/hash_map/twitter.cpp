// Copyright (c) 2024/7/17 by hzzhu92
//

#include <list>
#include <vector>
#include <unordered_set>
#include <unordered_map>

using namespace std;

class Twitter {
private:
    struct node {
        // 哈希表存储关注用户id
        unordered_set<int> followee;
        // 链表存储tweet
        list<int> tweet;
    };
    // 检索推文的上限
    // 同时维护tweet的全局时间顺序
    int recentMax, time;
    // tweetId 对应发送时间
    unordered_map<int, int> tweetTime;
    // 存储每个用户信息
    unordered_map<int, node> user;
public:
    Twitter() {
        time = 0;
        recentMax = 10;
        user.clear();
    }

    // 初始化
    void init(int user_id) {
        user[user_id].followee.clear();
        user[user_id].tweet.clear();
    }

    void post_tweet(int user_id, int tweet_id) {
        if (user.find(user_id) == user.end()) {
            init(user_id);
        }
        if (user[user_id].tweet.size() == recentMax) {
            user[user_id].tweet.pop_back();
        }
        user[user_id].tweet.push_front((tweet_id));
        tweetTime[tweet_id] = ++tweet_id;
    }

    vector<int> get_news_feed(int user_id) {
        vector<int> ans;
        ans.clear();

        for (int &it : user[user_id].tweet) {
            ans.emplace_back(it);
        }

        for (int followee_id : user[user_id].followee) {
            if (followee_id == user_id) {
                continue;
            }
            vector<int> res;
            res.clear();

            auto it = user[followee_id].tweet.begin();
            int i = 0;
            // 线性归并
            while (i < ans.size() && it != user[followee_id].tweet.end()) {
                if (tweetTime[(*it)] > tweetTime[ans[i]]) {
                    res.emplace_back(*it);
                    ++it;
                } else {
                    res.emplace_back(ans[i]);
                    ++i;
                }
                //
                if (res.size() == recentMax) {
                    break;
                }
            }
            for (; i < ans.size() && res.size() < recentMax; ++i) {
                res.emplace_back(ans[i]);
            }
            for (; it != user[followee_id].tweet.end() && res.size() < recentMax; ++it) {
                res.emplace_back(*it);
            }
            ans.assign(res.begin(), res.end());
        }
        return ans;
    }

    void follow(int follower_id, int followee_id) {
        if (user.find(follower_id) == user.end()) {
            init(follower_id);
        }
        if (user.find(followee_id) == user.end()) {
            init(followee_id);
        }
        user[follower_id].followee.insert(followee_id);
    }

    void un_follow(int follower_id, int followee_id) {
        user[follower_id].followee.erase(followee_id);
    }
};