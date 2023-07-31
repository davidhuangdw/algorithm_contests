fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    fun minIndex(a: List<Int>, st: Int = 0): Int {
        val n = a.size
        var i = st % n
        var j = st % n
        var (ci, cj) = 0 to 0
        while (ci < n && cj < n) {
            if (i == j) {
                j = (j+1)% n
                cj++
                continue
            }
            var k = 0
            while (k < n && a[(i + k) % n] == a[(j + k) % n]) {
                k++
            }
            if (k== n || a[(i + k) % n] <= a[(j + k) % n]) {
                cj += k+1
                j = (j+k+1) % n
            } else {
                ci += k+1
                i = (i+k+1) % n
            }
        }
        return if(ci < n) i else j
    }

//    for(arr in listOf(listOf(3,2, 1, 2, 1, 1, 2), listOf(2, 1, 1, 2, 1, 1))){
//        val i = minIndex(arr)
//        println("${arr}: ${i}, ${arr.subList(i, arr.size) + arr.subList(0, i)}")
//    }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, K) = input_ints()
        val A = input_ints()
        val B = input_ints()
        fun equal(i: Int, j: Int) = (0 until N).all { A[(i + it) % N] == B[(j + it) % N] }
        fun check(): Boolean {
            var (i, j) = minIndex(A) to minIndex(B)
            if(!equal(i, j)) return false

            if (N == 1) return true
            if (N == 2)
                return A[0] == B[K % 2]
            if(K > 1) return true
            if(K == 0) return i == j
            if(i == j)
                i = minIndex(A, i+1)
            return i != j
        }
        val r = if(check()) "YES" else "NO"
        println("Case #${it}: $r")
    }
}
