package google_kickstart.`2021`.C
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (G,) = input_longs()
        val mul = G*2
        var l = 1L
        var count = 0
        while(l*l < mul) {
            if (mul % l == 0L) {
                val y = mul / l
                if ((l + y) and 1L > 0L)
                    count += 1
            }
            l += 1
        }
        println("Case #${it}: $count")
    }
}
