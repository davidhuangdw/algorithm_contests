package google_kickstart.`2020`.F
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map{it.toInt()}
    fun input_longs() = split_input().map{it.toLong()}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (n, k) = input_ints()

        val ranges = (1..n).map{
            val (l, r) = input_ints()
            l to r
        }.sortedBy { it.first }

        var (cnt, ed) = 0 to 0
        for((x, y) in ranges){
            if(y <= ed) continue
            val st = maxOf(x, ed)
            val c = (y-st+k-1)/k
            cnt += c
            ed = st + c*k
        }
        println("Case #${it}: ${cnt}")
    }
}
