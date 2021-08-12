package google_kickstart.`2019`.F

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, S) = input_ints()
        val count = hashMapOf<Long, Int>()
        val skills = (0 until N).map {
            val line = input_ints()
            val sk = line.subList(1, line.size).sorted()

            var id = 0L
            for (v in sk)
                id = (id shl 10) + v
            count[id] = count.getOrDefault(id, 0) + 1

            sk
        }

        var sum = 0L
        for (sk in skills) {
            val c = sk.size
            sum += N
            for (bits in 0 until (1 shl c)) {
                var sub_id = 0L
                for (i in 0 until c)
                    if (bits shr i and 1 == 1)
                        sub_id = (sub_id shl 10) + sk[i]
                sum -= count.getOrDefault(sub_id, 0)
            }
        }

        println("Case #${it}: $sum")
    }
}
