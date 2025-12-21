package leetcode

import org.junit.jupiter.api.Assertions.assertArrayEquals
import org.junit.jupiter.api.Test

class RemoveInvalidParentheses301 {
    fun removeInvalidParentheses(s: String): List<String> {
        val n = s.length
        var total = 0
        var extra = 0
        for (ch in s) {
            if (ch == '(') {
                extra++
            } else if (ch == ')' && extra > 0) {
                total++
                extra--
            }
        }

        val remainL = IntArray(n+1)
        val remainR = IntArray(n+1)
        for(i in n-1 downTo 0){
            remainL[i] = remainL[i+1]
            remainR[i] = remainR[i+1]
            if(s[i] == '('){ remainL[i]++}
            if(s[i] == ')'){ remainR[i]++}
        }

        val ans = mutableListOf<String>()
        val cur = mutableListOf<Char>()
        fun dfs(i: Int, extraL: Int, remainPair: Int) {
            if (extraL + remainL[i] < remainPair || remainR[i] < remainPair) {
                return
            }

            if (i >= n) {
                if (remainPair == 0) {
                    ans.add(cur.joinToString(""))
                }
                return
            }

            val ch = s[i]
            if (ch in "()") {
                if(cur.isEmpty() || cur.last() != ch) {
                    dfs(i + 1, extraL, remainPair)
                }
                if (ch == '(') {
                    if (extraL < remainPair) {
                        cur.add(ch)
                        dfs(i + 1, extraL + 1, remainPair)
                        cur.removeLast()
                    }
                } else if (extraL > 0) {
                    cur.add(ch)
                    dfs(i + 1, extraL - 1, remainPair-1)
                    cur.removeLast()
                }
            } else {
                cur.add(ch)
                dfs(i + 1, extraL, remainPair)
                cur.removeLast()
            }
        }
        dfs(0, 0, total)
        return ans.toList()
    }


    @Test
    fun test1() {
        assertArrayEquals(removeInvalidParentheses("()())()").sorted().toTypedArray(), arrayOf("(())()", "()()()").sorted().toTypedArray())
        assertArrayEquals(removeInvalidParentheses("(a)())()").sorted().toTypedArray(), arrayOf("(a())()", "(a)()()").sorted().toTypedArray())
        assertArrayEquals(removeInvalidParentheses(")(").sorted().toTypedArray(), arrayOf("").sorted().toTypedArray())
    }
}
