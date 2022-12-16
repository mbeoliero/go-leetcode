/*
 * @lc app=leetcode.cn id=1494 lang=golang
 *
 * [1494] 并行课程 II
 */

// @lc code=start
func minNumberOfSemesters(n int, relations [][]int, k int) int {
	dp := make([]int, 1<<n)
	preReq := make(map[int]int)
	for _, val := range relations {
		pre := val[0] - 1
		after := val[1] - 1
		// pre[1] = 0110 表示1的先修课程是2和3
		preReq[after] |= 1 << pre
	}

	for state := 0; state < (1 << n); state++ {
		dp[state] = 2 << 32
	}

	// cnt 表示 i 中包含 1 的个数，即需要上多少课程
	cnt := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		cnt[i] = cnt[i>>1] + (i & 1)
	}

	// 0不需要上任何课程
	dp[0] = 0

	// taken表示已经上过的课程，taken = 0111表示已经上了1、2、3
	for taken := 0; taken < 1<<n; taken++ {
		if dp[taken] > n {
			// 将 n 门课程分开上时为n
			continue
		}

		// 当前可选课程放到cur中
		cur := 0
		for j := 0; j < n; j++ {
			// 表示 j 这门课还没在 taken 上过
			if (taken & (1 << j)) == 0 {
				// 表示 j 的先修课程都上完了
				if (preReq[j] & taken) == preReq[j] {
					// 把 j 加入到可选课程中来
					cur |= 1 << j
				}
			}
		}

		// 枚举 cur, cur = 111,它的子mask集合就是{111, 110, 101 011, 100, 010, 001}
		for subMask := cur; subMask > 0; subMask = (subMask - 1) & cur {
			if cnt[subMask] <= k {
				dp[taken|subMask] = min(dp[taken|subMask], dp[taken]+1)
			}
		}
	}

	return dp[(1<<n)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func minNumberOfSemesters(n int, dependencies [][]int, k int) int {
	//size of dp is 1<<n, start state is 000...000, end state is 111...111
	dp := make([]int, 1<<n)
	preReq := make(map[int]int)    //dict of state : prerequisite
	preCourse := make(map[int]int) //dict of course : pre-courses

	//calculate pre-courses of each course and maintain it as bitmask
	for _, val := range dependencies {
		//convert idx starting from 1 to 0
		//as the bitmask starts from 0
		preCourse[val[1]-1] += 1 << (val[0] - 1)
	}

	//calculate prerequisite
	for state := 0; state < (1 << n); state++ {
		dp[state] = 2 << 32
		preReq[state] = 0
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 { //the i-th digit is 1 means that course is taken in this state
				preReq[state] |= preCourse[i] //so this state's preRequisite should contain that course's pre-courses
			}
		}
	}
	dp[0] = 0 //need 0 semesters to complete 0 courses

	//iterate state from 000...001 to 111...111
	for state := 0; state < (1 << n); state++ {
		//preState possible values are from state to 0, however we don't have to iterate every single state
		//the reason is preState just need to be the subset of state
		//(preState - 1) & state is the logic to calculate state's subset
		preState := state
		for preState > 0 {
			preState = (preState - 1) & state

			//criteria 2, preState is superset of state's prerequisite
			//criteria 3, calculate number of 1s and compare with k
			if preReq[state] == (preState&preReq[state]) && courseDiff(state, preState) <= k {
				dp[state] = min(dp[state], dp[preState]+1)
			}
		}
	}

	return dp[(1<<n)-1] //return dp[111...111]
}

func courseDiff(a, b int) int {
	res := bits.OnesCount(uint(a)) - bits.OnesCount(uint(b))
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}