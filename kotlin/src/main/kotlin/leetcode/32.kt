package leetcode

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class LongestValidParentheses32 {

    fun longestValidParentheses(s: String): Int {  //2 pass
        // right point == 0
        var (mx, df) = 0 to 0
        var (l, r) = 0 to 0
        for(ch in s){
            if(ch == '(') l += 1
            else r += 1
            if(r > l){
                l = 0
                r = 0
            }else if(r == l)
                mx = maxOf(mx, l+r)
        }

        l = 0
        r = 0
        for(ch in s.reversed()){
            if(ch == '(') l += 1
            else r += 1
            if(l > r){
                l = 0
                r = 0
            }else if(l == r)
                mx = maxOf(mx, l+r)
        }
        return mx
    }

    fun longestValidParentheses2(s: String): Int {  //stack
        val st = mutableListOf(-1)
        var mx = 0
        for((i, ch) in s.withIndex()){
            if(ch == '(')
                st.add(i)
            else if(st.size > 1){
                st.removeAt(st.size-1)
                mx = maxOf(mx, i - st[st.size-1])
            }else{
                st[0] = i
            }
        }
        return mx
    }

    fun longestValidParentheses1(s: String): Int {  //dp
        val n = s.length
        val dp = MutableList(n){0}
        for((i, ch) in s.withIndex()){
            if(ch == ')' && i-1 >= 0){
                val l = i - dp[i-1]-1
                if(l >= 0 && s[l] == '(')
                    dp[i] = i - l + 1 + (if(l-1 >= 0) dp[l-1] else 0)
            }
        }
        return dp.maxOrNull() ?: 0
    }

    @Test
    fun test1() {
        assertEquals(0, longestValidParentheses(""))
        assertEquals(6, longestValidParentheses("(()())))()"))
        assertEquals(8, longestValidParentheses("(()())()((()"))
        assertEquals(2, longestValidParentheses("()((()"))
        assertEquals(4, longestValidParentheses("((((()()(()()(()()(()"))
    }
}
