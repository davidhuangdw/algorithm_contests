package atcoder.a6

fun main(){
    fun inputInts() = readLine()!!.split(" ").map(String::toInt)

    val (H, W, C, Q) = inputInts()
    var (h, w) = H to W
    val doneH = hashSetOf<Int>()
    val doneW = hashSetOf<Int>()
    val cnt = LongArray(C+1)

    val queries = (1..Q).map{inputInts()}
    for(q in queries.reversed()){
        val (t, i, c) = q
        if(t == 1){
            if(doneH.contains(i)) continue
            cnt[c] += w + 0L
            h--
            doneH.add(i)
        }else{
            if(doneW.contains(i)) continue
            cnt[c] += h + 0L
            w--
            doneW.add(i)
        }
    }

    println(cnt.asList().subList(1, C+1).joinToString(" "))
}