package atcoder.dp

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (NN, C) = input_longs()
    val N = NN.toInt()
    val H = input_longs()
    val F = MutableList(N) { 0L }
    val que = ArrayDeque<Int>()
    que.add(0)

    fun X(j: Int): Long = H[j]
    fun Y(j: Int): Long = F[j] + H[j] * H[j]
    fun compute(j: Int, k: Long): Long = Y(j) - k * X(j)
    fun kInc(j: Int, k: Int, i: Int): Boolean {
        return (Y(i) - Y(k)) * (X(k) - X(j)) > (Y(k) - Y(j)) * (X(i) - X(k))
    }
    for (i in 1 until N) {
        val k = 2 * H[i]
        // popFirst
        while (que.size > 1) {
            val a = que.removeFirst()
            val b = que.first()
            if (compute(a, k) < compute(b, k)) {
                que.addFirst(a)
                break
            }
        }

        val j = que.first()
        F[i] = F[j] + (H[j] - H[i]) * (H[j] - H[i]) + C
        // enqueLast
        while(que.size > 1){
            val a = que.removeLast()
            val b = que.last()
            if(kInc(b, a, i)){
                que.addLast(a)
                break
            }
        }
        que.add(i)
    }
    println(F[N-1])
}
