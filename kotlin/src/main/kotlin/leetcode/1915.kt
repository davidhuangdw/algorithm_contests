package leetcode

class WonderfulSubstrings {
    fun wonderfulSubstrings(word: String): Long {
        val n = 'j' - 'a' + 1
        var cur = 0
        val count = hashMapOf<Int, Long>()
        var res = 0L

        count[cur] = 1
        for (ch in word) {
            cur = cur xor (1 shl (ch - 'a'))

            res += count.getOrDefault(cur, 0L)
            for (i in 0 until n) {
                cur = cur xor (1 shl i)
                res += count.getOrDefault(cur, 0L)
                cur = cur xor (1 shl i)
            }

            count[cur] = count.getOrDefault(cur, 0L) + 1
        }
        return res
    }

}